package schedule

import (
	"context"
	"licron.com/schedule/cmd"
	"licron.com/schedule/constans"
	"licron.com/schedule/types"
	"sync"
	"time"
)

// 单体应用
type OnceSchedule struct {
	BaseSchedule
	// 时间执行任务
	cronExpr []*types.CronExpr
}

// 初始化操作
func NewOnce(chanNum int) *OnceSchedule {
	return &OnceSchedule{
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
func (s *OnceSchedule) Run() {
	// 获取所有的消息处理
	go s.schedule()
	// 获取到全部的计划任务
	crons, _ := s.onceModel.GetAll()
	//解析全部的任务 获取到任务的下次执行时间
	for _, r := range crons {
		t := &types.Cron{
			ID:       r.ID,
			Name:     r.Name,
			Command:  r.Command,
			Desc:     r.Desc,
			Runat:    r.Runat,
			IsKiller: 0,
			Types:    constans.OnceType,
		}
		s.addCron(t)
	}

}

// 添加任务
func (s *OnceSchedule) addCron(c *types.Cron) {
	now := time.Now()
	// 计算当前时间距离开始时间
	seconds := c.Runat.Sub(now).Seconds()
	if seconds < 0 {
		seconds = 0
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	newCmd := cmd.NewCmd(ctx, cancelFunc)
	s.execCancel[c.ID] = newCmd
	newCmd.ExecOnce(c, time.Duration(seconds)*time.Second, s.execChan)
}

// 一次执行任务 无法修改  先杀死 修改 再启动
func (s *OnceSchedule) updateCron(c *types.Cron) bool {
	return true
}
