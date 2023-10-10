package utils

import (
	"GinCoBlog/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

// resultType 返回值方法
func resultType(code int, msg string, data interface{}) config.Response {
	return config.Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

// SuccessResult 成功响应
func SuccessResult(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, resultType(http.StatusOK, msg, data))
}

// FailResult 错误响应
func FailResult(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, resultType(http.StatusBadRequest, msg, nil))
}

// ServerErrorResult 服务器错误响应
func ServerErrorResult(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, resultType(http.StatusInternalServerError, "服务器错误！", nil))
}

// AuthorizationResult 权限错误响应
func AuthorizationResult(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, resultType(http.StatusUnauthorized, msg, nil))
}

// StrIsEmpty 判断字符串为空
func StrIsEmpty(str string) bool {
	return str == "" || len(str) == 0
}

// VerPhoneReg /** 验证电话号码格式
func VerPhoneReg(phone string) bool {
	phoneReg := regexp.MustCompile(config.PhoneReg)
	return !phoneReg.MatchString(phone)
}

// VerEmailReg /** 验证电话号码格式
func VerEmailReg(email string) bool {
	emailReg := regexp.MustCompile(config.EmailReg)
	return !emailReg.MatchString(email)
}

// generateVerificationCode /** 生成验证码
func generateVerificationCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	code := ""
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", r.Intn(10))
	}
	return code
}

func SendEmail(email string) string {
	// 生成验证码
	code := generateVerificationCode()
	// 创建消息
	m := gomail.NewMessage()
	// 设置发件地址和发件人
	m.SetAddressHeader("From", config.EmailConfig.EmailAddress, config.EmailConfig.EmailName)
	// 发送地址
	m.SetHeader("To", email)
	// 设置标题
	m.SetHeader("Subject", "验证码")
	// 设置内容
	m.SetBody("text/html", `
            <p>您好！</p>
            <p>您的验证码是：<strong style="color:orangered;">`+code+`</strong></p>
			<p>此验证码在 5 分钟内有效</p>
            <p>如果不是您本人操作，请无视此邮件</p>
        `)
	// 使用 smtp发送邮件
	s := gomail.NewDialer(config.EmailConfig.SmtpServer, config.EmailConfig.SmtpPort, config.EmailConfig.EmailAddress, config.EmailConfig.Password)

	if err := s.DialAndSend(m); err != nil {
		panic("发送失败！")
	}
	return code
}
