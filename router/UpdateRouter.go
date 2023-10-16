package routes

import (
	"GinCoBlog/controller"
	"github.com/gin-gonic/gin"
)

// UploadRouter 定义文件上传路由（二级）
func UploadRouter(router *gin.RouterGroup) {
	// 用户头像
	router.PUT("/userAvatar", controller.UserAvatar)
	// 文章图片
	router.PUT("/articleIcon", controller.ArticleIcon)
}
