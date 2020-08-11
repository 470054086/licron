package request

// CronRequest请求表达式
type CronRequest struct {
	Name    string `json:"name" binding:"required"`
	Exp     string `json:"exp" binding:"required"`
	Command string `json:"command" binding:"required"`
	Desc string `json:"desc"`
}
