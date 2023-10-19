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
	router.GET("/articleList", controller.ArticleListByType)
	// 通过 ID 获取文章详情
	router.GET("/articleInfo", controller.ArticleInfoById)
	// 添加文章
	router.POST("/addArticle", controller.AddArticle)
	// 更新阅读量
	router.POST("/updateRead", controller.UpdateRead)
	// 获取文章评论
	router.GET("/articleComment", controller.ArticleComment)
	// 添加文章评论
	router.POST("/addComment", controller.AddArticleComment)

}
