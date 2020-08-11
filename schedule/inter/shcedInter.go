package inter

import (
	"licron.com/schedule/types"
)

// 任务的通知接口方法
type Schedule interface {
	AddNotify(c *types.Deamon)
	DelNotify(c *types.Deamon)
	UpdateNotify(c *types.Deamon)
	KillNotify(c *types.Deamon)
}