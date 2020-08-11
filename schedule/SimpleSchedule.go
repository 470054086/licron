package schedule

import (
	"context"
	"github.com/gorhill/cronexpr"
	"licron.com/global"
	"licron.com/model"
	"licron.com/schedule/types"
	"sort"
	"time"
)

// 单体应用
type SimpleSchedule struct {
	// 任务执行完成的通知
	execChan chan *types.CronExcel
	// 任务杀死的通知
	execKillChan chan *types.CronExpr
	// 添加任务的通知
	addSchedule chan *types.Cron
	// 删除任务的通知  不杀死当前执行的任务 只是删除
	deleteSchedule chan *types.Cron
	// killer掉任务的通知 killer掉 直接将当前任务杀死 然后删除
	killSchedule chan *types.Cron
	// 修改任务
	updateSchedule chan *types.Cron
	// 当前正在运行的
	cronRun []*types.CronExpr
	// execCancel 通过id和命令状态绑定 用于杀死任务
	execCancel map[int]*ExecCommand
	// 正在执行的任务
	executionSchedule map[int]*types.Cron
}

var c model.Cron

// 初始化操作
func New(execNum, addNum, delNum, killNum, updateNum int) *SimpleSchedule {
	return &SimpleSchedule{
		execChan:       make(chan *types.CronExcel, execNum),
		execKillChan:   make(chan *types.CronExpr, execNum),
		addSchedule:    make(chan *types.Cron, addNum),
		deleteSchedule: make(chan *types.Cron, delNum),
		killSchedule:   make(chan *types.Cron, killNum),
		updateSchedule: make(chan *types.Cron, updateNum),
		execCancel:     make(map[int]*ExecCommand),
	}
}

// 启动Schedule
func (s *SimpleSchedule) Run() {
	// 获取所有的消息处理
	go s.schedule()

	// 获取到全部的计划任务
	crons, _ := c.GetAll()
	//解析全部的任务 获取到任务的下次执行时间
	for _, r := range crons {
		s.addCron(r)
	}
	// 进行计划任务的处理
	for {
		for _, r := range s.cronRun {
			now := time.Now()
			// 如果小于或者等于当前时间的话 就执行
			if r.Next.Before(now) || r.Next.Equal(now) {
				// 任务是否在执行
				if s.isExecution(r.Cron) {
					continue
				}
				// 绑定命令和任务的关系
				ctx, cancelFunc := context.WithCancel(context.Background())
				cmd := NewCmd(ctx, cancelFunc)
				s.execCancel[r.Cron.ID] = cmd
				// 执行计划任务
				go cmd.Exec(r, s.execChan)
				// 计算当前任务的下次执行时间
				nextNow := time.Now()
				parse := cronexpr.MustParse(r.Cron.Exp)
				next := parse.Next(nextNow)
				r.Next = next
			}
		}
		// 根据下次执行时间进行排序 睡眠最近的执行时间
		sort.Sort(types.CronSort(s.cronRun))
		time.Sleep(s.cronRun[0].Next.Sub(time.Now()))
	}
}

// 异步通知
func (s *SimpleSchedule) AddNotify(c *types.Cron) {
	s.addSchedule <- c
}

// 添加任务
func (s *SimpleSchedule) addCron(c *types.Cron) {
	parse := cronexpr.MustParse(c.Exp)
	now := time.Now()
	next := parse.Next(now)
	run := types.CronExpr{Next: next, Cron: c}
	s.cronRun = append(s.cronRun, &run)
}

func (s *SimpleSchedule) DelNotify(c *types.Cron) {
	s.deleteSchedule <- c
}

// 删除任务
func (s *SimpleSchedule) delCron(c *types.Cron) bool {
	for i, r := range s.cronRun {
		if r.Cron.ID == c.ID {
			s.cronRun = append(s.cronRun[:i], s.cronRun[i+1:]...)
			return true
		}
	}
	s.delExecution(c)
	return false
}

func (s *SimpleSchedule) UpdateNotify(c *types.Cron) {
	s.updateSchedule <- c
}

// 进行修改任务
func (s *SimpleSchedule) updateCron(c *types.Cron) bool {
	for _, r := range s.cronRun {
		if r.Cron.ID == c.ID {
			r.Cron = c
			return true
		}
	}
	return false
}

func (s *SimpleSchedule) KillNotify(c *types.Cron) {
	s.killSchedule <- c
}

// 杀死任务
func (s *SimpleSchedule) killCron(c *types.Cron) bool {
	for _, r := range s.cronRun {
		if r.Cron.ID == c.ID {
			if _, ok := s.execCancel[r.Cron.ID]; ok {
				s.execCancel[r.Cron.ID].CancelFunc(r, s.execKillChan)
				return true
			}
		}
	}
	s.delExecution(c)
	return false
}

// 杀死任务回调通知
func (s *SimpleSchedule) killCronBack(c *types.CronExpr) bool {
	s.delCron(c.Cron)
	// 接触绑定关系
	delete(s.execCancel, c.Cron.ID)
	// 接触正在执行中的任务
	s.delExecution(c.Cron)
	return true
}

// 判断任务是否在执行
func (s *SimpleSchedule) isExecution(c *types.Cron) bool {
	if _, ok := s.executionSchedule[c.ID]; ok {
		return true
	}
	return false
}

// 删除已经执行的任务
func (s *SimpleSchedule) delExecution(c *types.Cron) bool {
	delete(s.executionSchedule, c.ID)
	return true
}

func (s *SimpleSchedule) schedule() {
	for {
		select {
		// 收到任务执行完毕的通知
		case res := <-s.execChan:
			s.delExecution(res.Cron.Cron)
			if res.Err != nil {
				// 暂时记录到日志 方便
				global.G_LOG.Info(res)
			}
		case r := <-s.addSchedule:
			s.addCron(r)
		case r := <-s.deleteSchedule:
			s.delCron(r)
		case r := <-s.killSchedule:
			s.killCron(r)
		case r := <-s.updateSchedule:
			s.updateCron(r)
		case r := <-s.execKillChan:
			// 杀死任务成功
			s.killCronBack(r)
		}
	}
}
