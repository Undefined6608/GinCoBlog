package request

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
