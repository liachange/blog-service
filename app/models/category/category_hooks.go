package category

import (
	"blog-service/pkg/database"
	"gorm.io/gorm"
	"strconv"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用 修改分类级别
func (category *Category) BeforeSave(tx *gorm.DB) (err error) {
	if category.ParentID != "0" {
		var level int64
		database.DB.Model(&category).Where("id=?", category.ParentID).Pluck("level", &level)
		category.Level = strconv.FormatInt(level+1, 10)
	}
	return
}
