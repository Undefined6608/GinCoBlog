package entity

import "time"

// SysUser 用户表结构
type SysUser struct {
	UID           uint32    `json:"uid" gorm:"primaryKey;autoIncrement;not null"`                           // 主键 用户ID
	UserName      string    `json:"user_name" gorm:"not null;type:varchar(255)"`                            // 用户名
	Phone         string    `json:"phone" gorm:"not null;unique;size:11"`                                   // 电话号码 不允许为空 唯一
	HeadSculpture string    `json:"head_sculpture" gorm:"default:'http://39.101.72.168:81/image/icon.jpg'"` // 头像 默认值
	Password      string    `json:"password" gorm:"type:text;not null"`                                     // 密码 默认值
	Email         string    `json:"email" gorm:"not null;unique"`                                           // 邮箱 不允许为空 唯一
	Available     bool      `json:"available" gorm:"not null;default:0"`                                    // 是否注销 不允许为空 默认值
	Limit         uint8     `json:"limit" gorm:"not null;default:2"`                                        // 权限 不允许为空 默认值
	Integral      uint32    `json:"integral" gorm:"not null;default:0"`                                     //用户积分
	Member        uint8     `json:"member" gorm:"not null;default:2"`                                       // 会员: 0 超级会员、1 会员、2 非会员
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time;type:datetime(0);autoUpdateTime"`  // 创建时间
	UpdateTime    time.Time `json:"update_time" gorm:"column:update_time;type:datetime(0);autoUpdateTime"`  // 更新时间
	UUID          string    `json:"uuid" gorm:"not null;type:varchar(36)"`                                  // 用户身份
}

// TableName /** 复写默认方法，设置表名
func (SysUser) TableName() string {
	return "sys_user"
}
