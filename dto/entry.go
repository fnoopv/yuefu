package dto

type PaginationParam struct {
	Page     int `json:"page" form:"page" binding:"required,min=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"required,min=10,max=100"`
}

type PaginationDataEntry struct {
	List     any   `json:"list,omitempty"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
}
