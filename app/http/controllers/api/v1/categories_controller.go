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

func (ctrl *CategoriesController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := category.Paginate(c, 15)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
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

func (ctrl *CategoriesController) Update(c *gin.Context) {
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategoryUpdate); !ok {
		return
	}
	categoryModel.Name = request.Name
	categoryModel.State = request.State
	categoryModel.Sort = request.Sort
	categoryModel.Description = request.Description
	categoryModel.ImageUrl = request.ImageUrl
	categoryModel.ParentID = request.ParentID
	rowsAffected := categoryModel.Save()
	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c)
	}
}

func (ctrl *CategoriesController) Delete(c *gin.Context) {
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	rowsAffected := categoryModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}
	response.Abort500(c, "删除失败，请稍后再试~")
}
