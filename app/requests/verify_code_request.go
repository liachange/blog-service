package requests

import (
	"blog-service/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	Phone         string `json:"phone,omitempty" valid:"phone"`
}

// VerifyCodePhone 验证表单，返回长度等于零即通过
func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}
	// 2. 定制错误消息
	message := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称为 phone",
			"digits:手机号长度必须为11位数字",
		},
		"captcha_id": []string{
			"required:图片验证码ID为必填项，参数名称为 captcha_id",
		},
		"captcha_answer": []string{
			"required:图片验证码的答案为必填项，参数名称为 captcha_answer",
			"digits:图片验证码答案长度必须为 6 数字",
		},
	}
	errs := validate(data, rules, message)
	// 图片验证码
	_data := data.(*VerifyCodePhoneRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
	return errs
}

type VerifyCodeEmailRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	Email         string `json:"email,omitempty" valid:"email"`
}

// VerifyCodeEmail 验证表单，返回长度等于零即通过
func VerifyCodeEmail(data interface{}, c *gin.Context) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
		"email":          []string{"required", "email", "min:4", "max:30"},
	}
	// 2. 定制错误消息
	message := govalidator.MapData{
		"captcha_id": []string{
			"required:图片验证码标识为必填项，参数名称为 captcha_id",
		},
		"captcha_answer": []string{
			"required:图片验证码答案为必填项，参数名称为 captcha_answer",
			"digits:图片验证码答案长度必须为 6位",
		},
		"email": []string{
			"required:邮箱为必填项，参数名称为email",
			"email:邮箱格式不正确",
			"min:邮箱长度需大于4",
			"max:邮箱长度需小于30",
		},
	}
	errs := validate(data, rules, message)
	_data := data.(*VerifyCodeEmailRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
	return errs
}
