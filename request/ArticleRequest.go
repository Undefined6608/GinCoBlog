package request

// ArticleListByTypeParam 根据类型 ID 查询的参数
type ArticleListByTypeParam struct {
	TypeId int32 `form:"type_id" binding:"required"`
}

// ArticleInfoByIdParam 根据 ID 获取文章详情
type ArticleInfoByIdParam struct {
	ArticleId int32 `form:"article_id" binding:"required"`
}

// AddArticleParam 添加文章参数
type AddArticleParam struct {
	TypeId  uint32 `json:"typeId" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Context string `json:"context" binding:"required"`
	Icon    string `json:"icon" binding:"required"`
}

// UpdateReadParam 更新阅读量
type UpdateReadParam struct {
	ArticleId int32 `json:"article_id" binding:"required"`
}

// ArticleCommentParam 文章评论参数
type ArticleCommentParam struct {
	ArticleId int32 `form:"article_id" binding:"required"`
}

// AddArticleCommentParam 添加文章评论
type AddArticleCommentParam struct {
	ArticleId uint32 `json:"article_id" binding:"required"`
	Context   string `json:"context" binding:"required"`
}
