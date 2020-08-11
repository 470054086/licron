package schedule

import (
	"context"
	"fmt"
	"licron.com/global"
	"licron.com/model"
	"licron.com/schedule/types"
)

// 单体应用
type DaemonSchedule struct {
	// 任务执行完成的通知
	execChan chan *types.CronDeamonExcel
	// 任务杀死的通知
	execKillChan chan *types.Deamon
	// 添加任务的通知
	addSchedule chan *types.Deamon
	// 删除任务的通知  不杀死当前执行的任务 只是删除
	deleteSchedule chan *types.Deamon
	// killer掉任务的通知 killer掉 直接将当前任务杀死 然后删除
	killSchedule chan *types.Deamon
	// 修改任务
	updateSchedule chan *types.Deamon
	// execCancel 通过id和命令状态绑定 用于杀死任务
	execCancel map[int]*ExecCommand
	// 正在执行的任务
	executionSchedule map[int]*types.Deamon
	// 所有的任务
	crons []*types.Deamon
}

var deamon model.Daemon

// 初始化操作
func DaemonNew(execNum, addNum, delNum, killNum, updateNum int) *DaemonSchedule {
	return &DaemonSchedule{
		execChan:          make(chan *types.CronDeamonExcel, execNum),
		execKillChan:      make(chan *types.Deamon, execNum),
		addSchedule:       make(chan *types.Deamon, addNum),
		deleteSchedule:    make(chan *types.Deamon, delNum),
		killSchedule:      make(chan *types.Deamon, killNum),
		updateSchedule:    make(chan *types.Deamon, updateNum),
		execCancel:        make(map[int]*ExecCommand),
		executionSchedule: make(map[int]*types.Deamon),
	}
}

// 启动Schedule
func (s *DaemonSchedule) Run() {
	// 获取所有的消息处理
	go s.schedule()
	// 获取到全部的常驻内存服务
	crons, _ := deamon.GetAll()
	s.crons = crons
	// 循环全部的变量 为每个任务启动一个goruntime
	for _, r := range crons {
		if ok := s.isExecution(r); ok {
			continue
		}
		// 将每个任务进行绑定
		// 绑定命令和任务的关系
		ctx, cancelFunc := context.WithCancel(context.Background())
		cmd := NewCmd(ctx, cancelFunc)
		s.execCancel[r.ID] = cmd
		// 添加任务正在执行
		s.addExecution(r)
		// 执行计划任务
		go cmd.ExecDaemon(r, s.execChan)

	}

}

// 异步通知
func (s *DaemonSchedule) AddNotify(c *types.Deamon) {
	s.addSchedule <- c
}

// 添加任务
func (s *DaemonSchedule) addCron(c *types.Deamon) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cmd := NewCmd(ctx, cancelFunc)
	s.execCancel[c.ID] = cmd
	// 添加任务正在执行
	s.addExecution(c)
	// 执行计划任务
	go cmd.ExecDaemon(c, s.execChan)
	s.crons = append(s.crons, c)
}

func (s *DaemonSchedule) DelNotify(c *types.Deamon) {
	s.deleteSchedule <- c
}

// 删除任务
func (s *DaemonSchedule) delCron(c *types.Deamon) bool {

	return false
}

func (s *DaemonSchedule) UpdateNotify(c *types.Deamon) {
	s.updateSchedule <- c
}

// 进行修改任务
func (s *DaemonSchedule) updateCron(c *types.Deamon) bool {

	return false
}

func (s *DaemonSchedule) KillNotify(c *types.Deamon) {
	s.killSchedule <- c
}

// 杀死任务
func (s *DaemonSchedule) killCron(c *types.Deamon) bool {
	for _, r := range s.crons {
		if r.ID == c.ID {
			if _, ok := s.execCancel[r.ID]; ok {
				global.G_LOG.Info(fmt.Sprintf("任务ID%d,任务名称为%s,进行杀死", r.ID, r.Name))
				s.execCancel[r.ID].DeamonCancelFunc(r, s.execKillChan)
				return true
			}
		}
	}
	return false
}

// 杀死任务回调通知
func (s *DaemonSchedule) killCronBack(c *types.CronExpr) bool {

	return true
}

// 判断任务是否在执行
func (s *DaemonSchedule) addExecution(c *types.Deamon) bool {
	if _, ok := s.executionSchedule[c.ID]; ok {
		return false
	}
	s.executionSchedule[c.ID] = c
	return false
}

// 判断任务是否在执行
func (s *DaemonSchedule) isExecution(c *types.Deamon) bool {
	if _, ok := s.executionSchedule[c.ID]; ok {
		return true
	}
	return false
}

// 删除已经执行的任务
func (s *DaemonSchedule) delExecution(c *types.Deamon) bool {
	delete(s.executionSchedule, c.ID)
	return true
}

func (s *DaemonSchedule) schedule() {
	for {
		select {
		// 收到任务执行完毕的通知
		case res := <-s.execChan:
			s.delExecution(res.Cron)
			if res.Err != nil {
				// 暂时记录到日志 方便
				global.G_LOG.Info(res)
			}
		case r := <-s.addSchedule:
			s.addCron(r)
		case r := <-s.deleteSchedule:
			s.delCron(r)
		case r := <-s.killSchedule:
			global.G_LOG.Info("接受到killer任务")
			s.killCron(r)
		case r := <-s.updateSchedule:
			s.updateCron(r)
		case r := <-s.execKillChan:
			// 杀死任务成功
			fmt.Println(r)
			//s.killCronBack(r)
		}
	}
}
