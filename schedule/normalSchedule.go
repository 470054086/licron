package schedule

import (
	"context"
	"github.com/gorhill/cronexpr"
	"licron.com/schedule/cmd"
	"licron.com/schedule/constans"
	"licron.com/schedule/types"
	"sort"
	"sync"
	"time"
)

// 单体应用
type SimpleSchedule struct {
	BaseSchedule
	// 时间执行任务
	cronExpr []*types.CronExpr
}

// 初始化操作
func NewSimple(chanNum int) *SimpleSchedule {
	return &SimpleSchedule{
		BaseSchedule: BaseSchedule{
			execChan:          make(chan *types.CronExcel, chanNum),
			execKillChan:      make(chan *types.Cron, chanNum),
			addSchedule:       make(chan *types.Cron, chanNum),
			killSchedule:      make(chan *types.Cron, chanNum),
			updateSchedule:    make(chan *types.Cron, chanNum),
			execCancel:        make(map[int]*cmd.ExecCommand),
			executionSchedule: make(map[int]*types.Cron),
			lock:              &sync.RWMutex{},
		},
	}
}

// 启动Schedule
func (s *SimpleSchedule) Run() {
	// 获取所有的消息处理
	go s.schedule()
	// 获取到全部的计划任务
	crons, _ := s.cronModel.GetAll()
	//解析全部的任务 获取到任务的下次执行时间
	for _, r := range crons {
		t := &types.Cron{
			ID:       r.ID,
			Name:     r.Name,
			Exp:      r.Exp,
			Command:  r.Command,
			Desc:     r.Desc,
			IsKiller: 0,
			Types:    constans.NormalType,
		}
		s.addCron(t)
	}
	// 进行计划任务的处理
	for {
		for _, r := range s.cronExpr {
			now := time.Now()
			// 如果小于或者等于当前时间的话 就执行
			if r.Next.Before(now) || r.Next.Equal(now) {
				// 任务是否在执行
				if s.isExecution(r.Cron) {
					continue
				}
				// 绑定命令和任务的关系
				ctx, cancelFunc := context.WithCancel(context.Background())
				cmd := cmd.NewCmd(ctx, cancelFunc)
				s.execCancel[r.Cron.ID] = cmd
				// 执行计划任务
				s.addExecution(r.Cron)
				go cmd.Exec(r.Cron, s.execChan)
				// 计算当前任务的下次执行时间
				nextNow := time.Now()
				parse := cronexpr.MustParse(r.Cron.Exp)
				next := parse.Next(nextNow)
				r.Next = next
			}
		}
		// 根据下次执行时间进行排序 睡眠最近的执行时间
		if s.cronExpr != nil {
			sort.Sort(types.CronSort(s.cronExpr))
			time.Sleep(s.cronExpr[0].Next.Sub(time.Now()))
		}
		time.Sleep(time.Second * 1)
	}
}

// 添加任务
func (s *SimpleSchedule) addCron(c *types.Cron) {
	parse := cronexpr.MustParse(c.Exp)
	now := time.Now()
	next := parse.Next(now)
	run := &types.CronExpr{Next: next, Cron: c}
	s.cronExpr = append(s.cronExpr, run)
	s.cronRun = append(s.cronRun, c)
	// todo 这里为了效率 可以重新刷新下list 新加入了重新计划

}

// 修改任务
func (s *SimpleSchedule) updateCron(c *types.Cron) bool {
	for _, r := range s.cronRun {
		if r.ID == c.ID {
			r = c
		}
		return true
	}
	return false
}
