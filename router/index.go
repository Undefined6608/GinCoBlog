package routes

import "github.com/gin-gonic/gin"

// SetupRouterGroup 项目主路由（一级）
func SetupRouterGroup(router *gin.RouterGroup) {
	// 调取用户路由
	UserRouter(router.Group("/user"))
	// 调取上传路由
	UploadRouter(router.Group("/upload"))
	// 调取文章路由
	ArticleRouter(router.Group("/article"))
	// 调取反馈路由
	FeedBackRouter(router.Group("/feedback"))
}
