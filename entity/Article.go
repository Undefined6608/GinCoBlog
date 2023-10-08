package entity

// Article 用户表结构
type Article struct {
	ID             int32  `json:"id" gorm:"primaryKey;autoIncrement;not null"` // 文章ID
	TypeId         int32  `json:"typeId" gorm:"not null"`                      // 文章类型ID
	UserId         int32  `json:"userId" gorm:"not null"`                      // 用户ID
	Title          string `json:"title" gorm:"not null;unique"`                // 文章标题
	Context        string `json:"context" gorm:"not null;type:text"`           // 文章内容
	ArticleVisible bool   `json:"article_visible" gorm:"not null;default:1"`   // 文章是否显示
	Read           int32  `json:"read" gorm:"not null;default:0"`              // 阅读量
	Date           string `json:"date" gorm:"not null;type:varchar(255)"`      // 上传时间
	Icon           string `json:"icon" gorm:"not null"`                        // 文章图像
}

// TableName /** 复写默认方法，设置表名
func (Article) TableName() string {
	return "article"
}
