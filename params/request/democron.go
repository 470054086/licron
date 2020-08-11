package request
type DeamonRequest struct {
	Name    string `json:"name" binding:"required"`
	Command string `json:"command" binding:"required"`
	Desc string `json:"desc"`
}
type DeamonKillRequest struct {
	ID int `json:"id" binding:"required"`
}