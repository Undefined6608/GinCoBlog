package request

// ArticleListByTypeParam 根据类型 ID 查询的参数
type ArticleListByTypeParam struct {
	TypeId int32 `json:"type_id" binding:"required"`
}
