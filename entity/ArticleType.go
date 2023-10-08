package entity

// ArticleType 文章类型表
type ArticleType struct {
	ID          uint32 `json:"id" gorm:"primaryKey;autoIncrement;not null"` // 文章类型ID
	RootID      uint32 `json:"root_id" gorm:"not null"`                     // 文章类型根ID
	TypeName    string `json:"type_name" gorm:"not null;unique"`            // 文章类型名称
	TypeVisible bool   `json:"type_visible" gorm:"not null;default:1"`      // 类型是否显示
	Order       uint8  `json:"order" gorm:"not null;unique"`                // 排序
	Picture     string `json:"picture" gorm:"not null"`                     // 类型图片
	AddStatus   bool   `json:"edit_status" gorm:"not null; default:1"`      // 是否可加入文章
}

// TableName /** 复写默认方法，设置表名
func (ArticleType) TableName() string {
	return "article_type"
}
