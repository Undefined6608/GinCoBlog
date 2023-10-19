package routes

import (
	"GinCoBlog/controller"
	"github.com/gin-gonic/gin"
)

// FeedBackRouter 定义反馈路由（二级）
func FeedBackRouter(router *gin.RouterGroup) {
	// 添加反馈
	router.POST("/addFeedback", controller.AddFeedBack)
}
