package request
type PageRequest struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
}
