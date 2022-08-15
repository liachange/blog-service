package requests

import (
	"blog-service/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type RestPasswordRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func RestByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required: 手机号为必填项，参数名称为 phone",
			"digits:手机号长度需为11位数字",
		},
		"verify_code": []string{
			"required:验证码为必填项，参数名称为 required",
			"digits:验证码长度需为6位",
		},
		"password": []string{
			"required:密码为必填项，参数名称为 password",
			"min:密码长度需大于6位",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*RestPasswordRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}
