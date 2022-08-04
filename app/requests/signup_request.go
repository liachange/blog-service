package requests

import (
	"blog-service/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExist(data any, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称为 phone",
			"digits:手机长度必须为11位数字",
		},
	}
	return validate(data, rules, messages)
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateSignupEmailExist(data any, c *gin.Context) map[string][]string {
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
	messages := govalidator.MapData{
		"email": []string{
			"required:邮箱不能为空，参数名称为 email",
			"min:邮箱长度须大于4",
			"max:邮箱长度须小于30",
			"email:邮箱格式不正确，请提供有效的邮箱地址",
		},
	}
	return validate(data, rules, messages)
}

// SignupUsingPhoneRequest 通过手机注册的请求信息
type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	Password        string `json:"password,omitempty" valid:"password"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
	Name            string `json:"name,omitempty" valid:"name"`
}

func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"name":             []string{"required", "between:3,20", "not_exists:users,phone"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}
	message := govalidator.MapData{
		"phone": []string{
			"required:手机号未必填项，参数名称为 phone",
			"digits:手机号长度为 11 位数字",
		},
		"name": []string{
			"required:用户名为必填项，参数名称为 name",
			"between:用户名长度需在3~20之间",
		},
		"password": []string{
			"required:密码为必填项，参数名称为password",
			"min:密码长度必须为6位数字",
		},
		"password_confirm": []string{
			"required:确认密码为必填项，参数名称为 password_confirm",
		},
	}
	errs := validate(data, rules, message)
	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}
