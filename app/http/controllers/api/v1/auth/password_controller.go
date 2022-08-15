package auth

import (
	v1 "blog-service/app/http/controllers/api/v1"
	"blog-service/app/models/user"
	"blog-service/app/requests"
	"blog-service/pkg/response"
	"github.com/gin-gonic/gin"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

func (pc *PasswordController) RestByPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.ResetPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.RestByPhone); !ok {
		return
	}
	// 2. 更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}

func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	// 1. 验证表单
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}
	// 2. 更新密码
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
