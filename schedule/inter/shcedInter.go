package inter

import (
	"licron.com/schedule/types"
)

// 任务的通知接口方法
type Schedule interface {
	AddNotify(c *types.Cron)
	UpdateNotify(c *types.Cron)
	KillNotify(c *types.Cron)
}
