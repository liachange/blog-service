package v1

import (
	"blog-service/app/models/category"
	"blog-service/app/requests"
	"blog-service/pkg/response"
	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) Store(c *gin.Context) {
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}
	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
		Sort:        request.Sort,
		State:       request.State,
		ParentID:    request.ParentID,
		ImageUrl:    request.ImageUrl,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}
