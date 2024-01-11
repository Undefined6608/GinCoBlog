package request

import "time"

// ArticleCommentResponse 文章评论返回参数
type ArticleCommentResponse struct {
	ID            uint32 `json:"id"`             // 评论ID
	ArticleID     uint32 `json:"article_id"`     // 文章ID
	Context       string `json:"context"`        // 评论内容
	UserName      string `json:"user_name"`      // 用户名
	HeadSculpture string `json:"head_sculpture"` // 头像 默认值
	Integral      uint32 `json:"integral"`       // 积分
	Member        uint8  `json:"member"`         // 会员
}

type ArticleListResponse struct {
	ID             uint32    `json:"id"`
	TypeId         uint32    `json:"typeId"`
	UserId         uint32    `json:"userId"`
	Title          string    `json:"title"`
	Context        string    `json:"context"`
	ArticleVisible bool      `json:"article_visible"`
	Read           uint32    `json:"read"`
	Date           time.Time `json:"date"`
	Icon           string    `json:"icon"`
	UserName       string    `json:"user_name"`
	HeadSculpture  string    `json:"head_sculpture"`
	Integral       uint32    `json:"integral"`
	Member         uint8     `json:"member"`
	UUID           string    `json:"uuid"`
}
