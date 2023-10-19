package request

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
