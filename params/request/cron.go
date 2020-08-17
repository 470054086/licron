package request

// CronRequest请求表达式
type CronRequest struct {
	Name    string `json:"name" binding:"required"`
	Exp     string `json:"exp" binding:"required"`
	Command string `json:"command" binding:"required"`
	Desc    string `json:"desc"`
}

// 查询的参数
type CronListRequest struct {
	PageRequest
}

// 定义返回格式
type CronListResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Exp       string    `json:"exp"`
	Command   string    `json:"command"`
	Desc      string    `json:"desc"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsKiller  int       `json:"is_killer"`
}
