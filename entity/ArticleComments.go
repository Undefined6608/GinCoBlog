package entity

// ArticleComments 文章评论表
type ArticleComments struct {
	ID        uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null"` // 评论ID
	ArticleID uint32 `json:"article" gorm:"not null"`                     // 文章ID
	UserID    uint32 `json:"user_id" gorm:"not null"`                     // 用户ID
	Context   string `json:"context" gorm:"not null;type:varchar(255)"`   // 评论内容
}

// TableName /** 复写默认方法，设置表名
func (ArticleComments) TableName() string {
	return "article_comments"
}
