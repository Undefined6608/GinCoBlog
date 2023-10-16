package request

import (
	"GinCoBlog/entity"
	"github.com/dgrijalva/jwt-go"
)

// UserNameOccupyParams 用户名查重参数
type UserNameOccupyParams struct {
	UserName string `json:"user_name" binding:"required,min=3,max=15"`
}

// PhoneOccupyParams 电话号码查重参数
type PhoneOccupyParams struct {
	Phone string `json:"phone" binding:"required,len=11"`
}

// EmailOccupyParams 邮箱查重参数
type EmailOccupyParams struct {
	Email string `json:"email" binding:"required"`
}

// SendEmailParams 邮箱查重参数
type SendEmailParams struct {
	Email string `json:"email" binding:"required"`
}

// RegisterParams 注册参数
type RegisterParams struct {
	UserName    string `json:"user_name" binding:"required,min=3,max=15"`
	Phone       string `json:"phone" binding:"required,len=11"`
	Email       string `json:"email" binding:"required"`
	EmailCode   string `json:"email_code" binding:"required,len=6"`
	Password    string `json:"password" binding:"required,len=32"`
	VerPassword string `json:"ver_password" binding:"required,len=32"`
}

// PhoneLoginParams 电话号码登录参数
type PhoneLoginParams struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required,len=32"`
}

// TokenParams 定义 Token 类型
type TokenParams struct {
	UserInfo           entity.SysUser // 用户信息
	jwt.StandardClaims                // token 配置
}

// EmailLoginParams 邮箱登录参数
type EmailLoginParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,len=32"`
}

// ForgotPasswordParams 忘记密码参数
type ForgotPasswordParams struct {
	Email       string `json:"email" binding:"required"`
	EmailCode   string `json:"email_code" binding:"required,len=6"`
	NewPassword string `json:"new_password" binding:"required,len=32"`
	VerPassword string `json:"ver_password" binding:"required,len=32"`
}

// ModifyUserInfoParams 修改用户信息参数
type ModifyUserInfoParams struct {
	UserName      string `json:"user_name" binding:"required,min=3,max=15"`
	Phone         string `json:"phone" binding:"required,len=11"`
	Email         string `json:"email" binding:"required"`
	HeadSculpture string `json:"head_sculpture" binding:"required"`
	EmailCode     string `json:"email_code" binding:"required,len=6"`
}

// ModifyPasswordParams 修改密码参数
type ModifyPasswordParams struct {
	OldPassword      string `json:"old_password" binding:"required,len=32"`
	NewPassword      string `json:"new_password" binding:"required,len=32"`
	VerifiedPassword string `json:"verified_password" binding:"required,len=32"`
}
