package entity

// FeedBack 反馈表
type FeedBack struct {
	ID      uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null"` // 评论ID
	UserID  uint32 `json:"user_id" gorm:"not null"`                     // 用户ID
	Context string `json:"context" gorm:"not null;type:varchar(255)"`   // 反馈内容
}

// TableName /** 复写默认方法，设置表名
func (FeedBack) TableName() string {
	return "feed_back"
}
