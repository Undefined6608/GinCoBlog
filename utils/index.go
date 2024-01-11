package utils

import (
	"GinCoBlog/config"
	"GinCoBlog/request"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
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
	c.JSON(http.StatusOK, resultType(http.StatusBadRequest, msg, nil))
}

// ServerErrorResult 服务器错误响应
func ServerErrorResult(c *gin.Context) {
	c.JSON(http.StatusOK, resultType(http.StatusInternalServerError, "服务器错误！", nil))
}

// AuthorizationResult 权限错误响应
func AuthorizationResult(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, resultType(http.StatusUnauthorized, msg, nil))
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

// SendEmail 发送邮箱
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

// CreateUUID 生成 uuid
func CreateUUID() string {
	uCode := uuid.NewV4()
	return uCode.String()
}

// EncryptionPassword 密码加密
func EncryptionPassword(pwd string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(pwd+config.Encryption.PrivateKey.Password), config.Encryption.Salt.Password)
	if err != nil {
		return ""
	}
	return string(password)
}

// ComparePassword 密码验证
func ComparePassword(hashPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd+config.Encryption.PrivateKey.Password))
	if err != nil {
		return false
	}
	return true
}

// GenerateToken 生成 Token
func GenerateToken(claims *request.TokenParams) string {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	claims.ExpiresAt = time.Now().Add(config.TokenEffectAge).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.TokenPrivateKey))
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		log.Panicln("Token生成异常")
		return ""
	}
	return sign
}

// IsContainArr Token URL过滤
func IsContainArr(noVerify []string, requestUrl string) bool {
	for _, str := range noVerify {
		if str == requestUrl {
			return true
		}
	}
	return false
}

// GetCacheUser 获取缓存里的token信息
func GetCacheUser(c *gin.Context) (error, *request.TokenParams) {
	// 获取用户信息
	user, _ := c.Get("user")
	// 判断用户信息是否存在
	if user == nil {
		return errors.New("登录状态错误，请重新登录"), nil
	}
	// 将user转化为 TokenParam类型
	tokenParam, ok := user.(*request.TokenParams)
	// 判断是否转化正确
	if !ok {
		return errors.New("登录状态错误，请重新登录"), nil
	}
	return nil, tokenParam
}

// GenerateFileName 生成文件名
func GenerateFileName(originalName string) string {
	// 提取文件后缀
	extension := filepath.Ext(originalName)
	// 生成 UUID 字符串
	uuidFilename := uuid.NewV4().String()
	// 返回拼接字符串
	return uuidFilename + extension
}

// IsAllowedImageType 定义允许上传的文件类型
func IsAllowedImageType(extension string) bool {
	// 获取允许的类型
	allowedImageTypes := config.Upload.ImgType
	// 判断是否符合该类型
	return contains(allowedImageTypes, extension)
}

// 判断是否允许上传
func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
