package cmd

import (
	"context"
	"fmt"
	"licron.com/global"
	"licron.com/schedule/types"
	"os/exec"
	"time"
)

type ExecCommand struct {
	ctx    context.Context
	cancel context.CancelFunc
	isKill bool
}

func NewCmd(ctx context.Context, cancelFunc context.CancelFunc) *ExecCommand {
	return &ExecCommand{
		ctx:    ctx,
		cancel: cancelFunc,
		isKill: true,
	}
}

// 执行任务 这里停止任务
func (e *ExecCommand) Exec(c *types.Cron, s chan *types.CronExcel) {
	//// todo 这里暂时为加锁 通过锁可以防止同一个任务 多次执行 稍后添加
	global.G_LOG.Info(fmt.Sprintf("定时任务任务名称为%s,开始执行时间为%s", c.Name, time.Now()))
	commandContext := exec.CommandContext(e.ctx, "base", "-c", c.Command)
	output, err := commandContext.Output()
	s <- &types.CronExcel{Output: output, Err: err, Cron: c}
}

// 执行常驻内存的服务
// todo 测试环境使用 线上的话 线上的话 不会走
func (e *ExecCommand) ExecDaemon(c *types.Cron, s chan *types.CronExcel) {
	// todo 手动模拟常驻内存
	//for {
	//	if !e.isKill {
	//		continue
	//	}
	//	global.G_LOG.Info(fmt.Sprintf("常驻内存任务名称为%s,开始执行时间为%s", c.Name, time.Now()))
	//	commandContext := exec.CommandContext(e.ctx, "D:/软件/Git/bin/bash.exe", "-c", c.Command)
	//	output, err := commandContext.Output()
	//	global.G_LOG.Info(fmt.Sprintf("任务的执行结果为%s,错误为%v", output, err))
	//	s <- &types.CronExcel{
	//		Cron:   c,
	//		Output: output,
	//		Err:    err,
	//	}
	//	e.isKill = true
	//	time.Sleep(time.Second * 10)
	//}
}

// 杀死进程
func (e *ExecCommand) CancelFunc(c *types.Cron, s chan *types.Cron) {
	e.cancel()
	s <- c
}
