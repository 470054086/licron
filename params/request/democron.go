package request
type DeamonRequest struct {
	Name    string `json:"name" binding:"required"`
	Command string `json:"command" binding:"required"`
	Desc string `json:"desc"`
}
type DeamonKillRequest struct {
	ID int `json:"id" binding:"required"`
}

// 查询的参数
type DeamonListRequest struct {
	PageRequest
}

// 定义返回格式
type DaemonListResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Command   string    `json:"command"`
	Desc      string    `json:"desc"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsKiller  int       `json:"is_killer"`
}