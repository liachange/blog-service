package auth

import (
	"blog-service/app/models/user"
	"blog-service/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
)

// Attempt 尝试登录
func Attempt(loginID, password string) (user.User, error) {
	userModel := user.GetByMulti(loginID)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账户不存在")
	}
	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码不正确")
	}
	return userModel, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号不存在")
	}
	return userModel, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	return userModel
}

// CurrentUserID 从 gin.context 中获取当前登录用户 ID
func CurrentUserID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
