package schedule

import (
	"fmt"
	"licron.com/global"
	"licron.com/model"
	"licron.com/schedule/cmd"
	"licron.com/schedule/constans"
	"licron.com/schedule/types"
	"sync"
)

type BaseSchedule struct {
	// 任务执行完成的通知
	execChan chan *types.CronExcel
	// 任务杀死的通知
	execKillChan chan *types.Cron
	// 添加任务的通知
	addSchedule chan *types.Cron
	// killer掉任务的通知 killer掉 直接将当前任务杀死 然后删除
	killSchedule chan *types.Cron
	// 修改任务
	updateSchedule chan *types.Cron
	// execCancel 通过id和命令状态绑定 用于杀死任务
	execCancel map[int]*cmd.ExecCommand
	// 正在执行的任务
	executionSchedule map[int]*types.Cron
	cronModel         model.Cron
	daemonModel       model.Daemon
	onceModel         model.OnceCron

	// 当前正在执行的任务
	cronRun []*types.Cron
	lock    *sync.RWMutex
}

// 启动Schedule
func (s *BaseSchedule) Run() {
}

// 异步通知
func (s *BaseSchedule) AddNotify(c *types.Cron) {
	s.addSchedule <- c
}

// 添加任务
func (s *BaseSchedule) addCron(c *types.Cron) {

}

func (s *BaseSchedule) UpdateNotify(c *types.Cron) {
	s.updateSchedule <- c
}

// 修改任务
func (s *BaseSchedule) updateCron(c *types.Cron) bool {
	return true
}

// 杀死任务通知
func (s *BaseSchedule) KillNotify(c *types.Cron) {
	s.killSchedule <- c
}

// 杀死任务
func (s *BaseSchedule) killCron(c *types.Cron) bool {
	for _, r := range s.cronRun {
		if r.ID == c.ID {
			if _, ok := s.execCancel[r.ID]; ok {
				global.G_LOG.Info(fmt.Sprintf("任务ID%d,任务名称为%s,进行杀死", r.ID, r.Name))
				s.execCancel[r.ID].CancelFunc(r, s.execKillChan)
				return true
			}
		}
	}
	return false
}

// 杀死任务回调通知
func (s *BaseSchedule) killCronBack(c *types.Cron) bool {
	global.G_LOG.Info(fmt.Sprintf("任务kill成功,任务ID为%d,任务名称为%s", c.ID, c.Name))
	// 取消正在运行
	s.delExecution(c)
	switch c.Types {
	case constans.NormalType:
		s.cronModel.Killer(c.ID)
	case constans.DaemonType:
		// 改变数据库状态
		s.daemonModel.Killer(c.ID)
	case constans.OnceType:
		s.onceModel.Killer(c.ID)
	}
	return true
}

// 判断任务是否在执行
func (s *BaseSchedule) addExecution(c *types.Cron) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.executionSchedule[c.ID]; ok {
		return false
	}
	s.executionSchedule[c.ID] = c
	return false
}

// 判断任务是否在执行
func (s *BaseSchedule) isExecution(c *types.Cron) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.executionSchedule[c.ID]; ok {
		return true
	}
	return false
}

// 删除已经执行的任务
func (s *BaseSchedule) delExecution(c *types.Cron) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.executionSchedule, c.ID)
	return true
}

func (s *BaseSchedule) schedule() {
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
		case r := <-s.killSchedule:
			global.G_LOG.Info("接受到killer任务")
			s.killCron(r)
		case r := <-s.updateSchedule:
			s.updateCron(r)
		case r := <-s.execKillChan:
			s.killCronBack(r)
		}
	}
}
