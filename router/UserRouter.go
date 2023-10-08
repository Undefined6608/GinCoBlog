package routes

import (
	"GinCoBlog/controller"
	"github.com/gin-gonic/gin"
)

// UserRouter 定义用户路由（二级）
func UserRouter(router *gin.RouterGroup) {
	// 用户相关接口
	router.GET("/", controller.HelloUser)
}
