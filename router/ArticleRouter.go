package routes

import (
	"GinCoBlog/controller"
	"github.com/gin-gonic/gin"
)

// ArticleRouter  定义文章路由（二级）
func ArticleRouter(router *gin.RouterGroup) {
	// 获取文章类型列表
	router.GET("/articleType", controller.ArticleType)
	// 通过类型 ID 获取文章列表
	router.POST("/articleList", controller.ArticleListByType)
}
