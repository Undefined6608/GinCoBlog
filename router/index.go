package routes

import "github.com/gin-gonic/gin"

// SetupRouterGroup 项目主路由（一级）
func SetupRouterGroup(router *gin.RouterGroup) {
	// 调取用户路由
	UserRouter(router.Group("/user"))
}
