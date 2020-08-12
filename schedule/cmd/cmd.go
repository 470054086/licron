package cmd

import (
	"context"
	"fmt"
	"licron.com/global"
	"licron.com/model"
	"licron.com/schedule/types"
	"os/exec"
	"time"
)

type ExecCommand struct {
	ctx       context.Context
	cancel    context.CancelFunc
	isKill    bool
	onceModel model.OnceCron
}

func NewCmd(ctx context.Context, cancelFunc context.CancelFunc) *ExecCommand {
	return &ExecCommand{
		ctx:       ctx,
		cancel:    cancelFunc,
		isKill:    true,
		onceModel: model.OnceCron{},
	}
}

// 执行任务 这里停止任务
func (e *ExecCommand) Exec(c *types.Cron, s chan *types.CronExcel) {
	if !e.isKill {
		return
	}
	global.G_LOG.Info(fmt.Sprintf("定时任务任务名称为%s,开始执行时间为%s", c.Name, time.Now()))
	commandContext := exec.CommandContext(e.ctx, "base", "-c", c.Command)
	output, err := commandContext.Output()
	s <- &types.CronExcel{Output: output, Err: err, Cron: c}
}

func (e *ExecCommand) ExecOnce(c *types.Cron, diff time.Duration, s chan *types.CronExcel) {
	fmt.Println(diff)
	time.AfterFunc(diff, func() {
		global.G_LOG.Info(fmt.Sprintf("任务名称为%s,开始执行时间为%s", c.Name, time.Now()))
		commandContext := exec.CommandContext(e.ctx, "D:/软件/Git/bin/bash.exe", "-c", c.Command)
		output, err := commandContext.Output()
		global.G_LOG.Info(fmt.Sprintf("任务的执行结果为%s,错误为%v", output, err))
		s <- &types.CronExcel{
			Cron:   c,
			Output: output,
			Err:    err,
		}
		e.isKill = true
		time.Sleep(time.Second * 10)
	})
}

// 杀死进程
func (e *ExecCommand) CancelFunc(c *types.Cron, s chan *types.Cron) {
	e.cancel()
	e.isKill = false
	s <- c
}
