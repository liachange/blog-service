package category

import "blog-service/pkg/database"

// Get 通过 ID 获取用户
func Get(idstr string) (categoryModel Category) {
	database.DB.Where("id", idstr).First(&categoryModel)
	return
}
