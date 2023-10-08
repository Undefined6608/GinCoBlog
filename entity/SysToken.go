package entity

type SysToken struct {
	ID     int32  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserId uint32 `json:"user_id" gorm:"not null"`
	Token  string `json:"token" gorm:"not null;type:varchar(255)"`
}

// TableName /** 复写默认方法，设置表名
func (SysToken) TableName() string {
	return "sys_token"
}
