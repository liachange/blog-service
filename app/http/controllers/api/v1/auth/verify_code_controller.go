package auth

import (
	v1 "blog-service/app/http/controllers/api/v1"
	"blog-service/app/requests"
	"blog-service/pkg/captcha"
	"blog-service/pkg/logger"
	"blog-service/pkg/response"
	"blog-service/pkg/verifycode"
	"github.com/gin-gonic/gin"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	//生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)
	// 返回给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// SendUsingPhone 发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}
	// 2. 发送 SMS
	if err := verifycode.NewVerifyCode().SendSMS(request.Phone); !err {
		response.Abort500(c, "短信发送失败")
	} else {
		response.Success(c)
	}
}

// SendUsingEmail 发送 Email 验证码
func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {
	//1.验证表单
	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}
	//2.发送邮件
	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(c, "邮件发送失败")
	} else {
		response.Success(c)
	}
}
