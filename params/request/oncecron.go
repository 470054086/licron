package request

import "time"

// CronRequest请求表达式
type OnceRequest struct {
	Name    string    `json:"name" binding:"required"`
	Command string    `json:"command" binding:"required"`
	Runat   time.Time `json:"runat" binding:"required" `
	Desc    string    `json:"desc"`
}

// 查询的参数
type OnceListRequest struct {
	PageRequest
}
