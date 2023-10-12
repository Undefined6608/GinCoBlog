package request

type UserResponse struct {
	Name          string `json:"name"`           // 用户名
	Phone         string `json:"phone"`          // 电话号码 不允许为空 唯一
	HeadSculpture string `json:"head_sculpture"` // 头像 默认值
	Email         string `json:"email"`          // 邮箱 不允许为空 唯一
	Limit         uint8  `json:"limit"`          // 权限 不允许为空 默认值
	Integral      uint32 `json:"integral"`       // 积分
	Member        uint8  `json:"member"`         // 会员
	CreateTime    string `json:"create_time"`    // 创建时间
	UpdateTime    string `json:"update_time"`    // 更新时间
	UUID          string `json:"uuid"`           // 用户证书
}
