package requests

import (
	"blog-service/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

// LoginByPhone 验证表单，返回长度等于零即通过
func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填写，参数名称为 phone",
			"digits:手机号长度必须为11位数字",
		},
		"verify_code": []string{
			"required:验证码答案为必填项，参数名称为 verify_code",
			"digits:验证码答案长度必须为 6 位数字",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}
