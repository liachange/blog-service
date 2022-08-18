package category

import (
	"blog-service/pkg/app"
	"blog-service/pkg/database"
	"blog-service/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// Get 通过 ID 获取用户
func Get(idstr string) (categoryModel Category) {
	database.DB.Where("id", idstr).First(&categoryModel)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Category{}),
		&categories,
		app.V1URL(database.TableName(&Category{})),
		perPage,
	)
	return
}
