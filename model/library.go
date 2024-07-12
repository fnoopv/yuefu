package model

type Library struct {
	ID        string `json:"id" gorm:"column:id;primaryKey"`
	Mode      string `json:"mode" gorm:"column:mode;not null"`
	Path      string `json:"path" gorm:"column:path;not null;unique"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (Library) TableName() string {
	return "libraries"
}
