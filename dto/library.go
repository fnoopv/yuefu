package dto

type Library struct {
	ID        string `json:"id"`
	Mode      string `json:"mode"`
	Path      string `json:"path"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type LibraryList struct {
	Mode string `json:"mode" form:"mode" binding:"required_with=import collection"`
	PaginationParam
}

type LibraryCreate struct {
	Mode string `json:"mode" binding:"required_with=import collection"`
	Path string `json:"path" binding:"required"`
}

type LibraryUpdate struct {
	ID   string `json:"id" binding:"required"`
	Path string `json:"path" binding:"required"`
}
