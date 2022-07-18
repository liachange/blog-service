package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExistRequest(data any, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	// 自定义验证出错时的提示
	message := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称为 phone",
			"digits:手机长度必须为11位数字",
		},
	}
	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      message,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateSignupEmailExistRequest(data any, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{
			"required",
			"min:4",
			"max:30",
			"email",
		},
	}
	// 自定义验证出错时的提示
	message := govalidator.MapData{
		"email": []string{
			"required:邮箱不能为空，参数名称为 email",
			"min:邮箱长度须大于4",
			"max:邮箱长度须小于30",
			"email:邮箱格式不正确，请提供有效的邮箱地址",
		},
	}
	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      message,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
