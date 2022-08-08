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

type LoginByPasswordRequest struct {
	LoginID       string `json:"login_id,omitempty" valid:"login_id"`
	Password      string `json:"password,omitempty" valid:"password"`
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
}

// LoginByPassword 验证表单，返回长度等于零即通过
func LoginByPassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required"},
		"login_id":       []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"captcha_id": []string{
			"required:验证码 ID 为必填项，参数名称为 captcha_id",
		},
		"captcha_answer": []string{
			"required:验证码答案为必填,项参数名称为 captcha_answer",
		},
		"login_id": []string{
			"required:登录账户为必填项，参数名称为 login_id",
			"min:登录账户长度需大于3",
		},
		"password": []string{
			"required:密码为必填项，参数名称为 password",
		},
	}
	errs := validate(data, rules, messages)
	// 图片验证码
	_data := data.(*LoginByPasswordRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
	return errs
}
