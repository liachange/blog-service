package category

import (
	"blog-service/app/models"
	"blog-service/pkg/database"
	"time"
)

type Category struct {
	models.BaseModel
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	State       string    `json:"state,omitempty"`
	ImageUrl    string    `json:"image_url,omitempty"`
	Sort        string    `json:"sort,omitempty"`
	ParentID    string    `json:"parent_id,omitempty" gorm:"default:0"`
	Level       string    `json:"level,omitempty" gorm:"default:1"`
	DeletedAt   time.Time `json:"-" gorm:"default:''"`
	models.CommonTimestampsField
}

func (category *Category) Create() {
	database.DB.Create(&category)
}
