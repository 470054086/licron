package schedule

import (
	"context"
	"licron.com/model"
	"licron.com/schedule/cmd"
	"licron.com/schedule/types"
	"licron.com/service"
	"sync"
)

// 单体应用
type DaemonSchedule struct {
	BaseSchedule
}

var (
	daemon        model.Daemon
	daemonService service.DeamonService
)

// 初始化操作
func DaemonNew(execNum, addNum, delNum, killNum, updateNum int) *DaemonSchedule {
	return &DaemonSchedule{
		BaseSchedule: BaseSchedule{
			execChan:             make(chan *types.CronExcel, execNum),
			execKillChan:         make(chan *types.Cron, execNum),
			addSchedule:          make(chan *types.Cron, addNum),
			killSchedule:         make(chan *types.Cron, killNum),
			updateSchedule:       make(chan *types.Cron, updateNum),
			execCancel:           make(map[int]*cmd.ExecCommand),
			executionSchedule:    make(map[int]*types.Cron),
			lock:                 &sync.RWMutex{},
		},
	}

}

// 启动Schedule
func (s *DaemonSchedule) Run() {
	// 获取所有的消息处理
	go s.schedule()
	// 获取到全部的常驻内存服务
	crons, _ := daemon.GetAll()
	for _, v := range crons {
		t := daemonService.TransFrom(v)
		s.cronRun = append(s.cronRun, t)
	}
	// 循环全部的变量 为每个任务启动一个goruntime
	for _, r := range s.cronRun {
		if ok := s.isExecution(r); ok {
			continue
		}
		// 将每个任务进行绑定
		// 绑定命令和任务的关系
		s.addCron(r)
	}
}

// 添加任务
func (s *DaemonSchedule) addCron(c *types.Cron) {
	// todo 记录context 到时候取消 goruntime
	ctx, cancelFunc := context.WithCancel(context.Background())
	newCmd := cmd.NewCmd(ctx, cancelFunc)
	s.execCancel[c.ID] = newCmd
	// 添加任务正在执行
	s.addExecution(c)
	// 执行计划任务
	go newCmd.ExecDaemon(c, s.execChan)
	s.cronRun = append(s.cronRun, c)

}

// 修改任务
// 常驻内存的话 不提供修改  先杀死 修改完成之后 再启动
func (s *DaemonSchedule) updateCron(c *types.Cron) bool {

	return false
}
