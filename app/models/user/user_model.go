package user

import "blog-service/app/models"

// User 用户模型
type User struct {
	models.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	models.CommonTimestampsField
}
