package request

// ArticleListByTypeParam 根据类型 ID 查询的参数
type ArticleListByTypeParam struct {
	TypeId int32 `json:"type_id" binding:"required"`
}

// ArticleInfoByIdParam 根据 ID 获取文章详情
type ArticleInfoByIdParam struct {
	ArticleId int32 `json:"article_id" binding:"required"`
}

// AddArticleParam 添加文章参数
type AddArticleParam struct {
	TypeId  uint32 `json:"typeId" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Context string `json:"context" binding:"required"`
	Icon    string `json:"icon" binding:"required"`
}

type UpdateReadParam struct {
	ArticleId int32 `json:"article_id" binding:"required"`
}
