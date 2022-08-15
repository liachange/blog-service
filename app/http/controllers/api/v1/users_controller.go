package v1

import (
	"blog-service/app/models/user"
	"blog-service/pkg/auth"
	"blog-service/pkg/response"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

//Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	data := user.All()
	response.Data(c, data)
}
