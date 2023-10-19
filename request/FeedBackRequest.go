package request

// AddFeedBackParam 添加反馈信息参数
type AddFeedBackParam struct {
	Context string `json:"" binding:"required"`
}
