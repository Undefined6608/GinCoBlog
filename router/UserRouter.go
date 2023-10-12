package routes

import (
	"GinCoBlog/controller"
	"github.com/gin-gonic/gin"
)

// UserRouter 定义用户路由（二级）
func UserRouter(router *gin.RouterGroup) {
	// 用户相关接口
	// 测试接口
	router.GET("/", controller.HelloUser)
	// 用户名查重
	router.GET("/userNameOccupy", controller.UserNameOccupy)
	// 电话号码查重
	router.GET("/phoneOccupy", controller.PhoneOccupy)
	// 邮箱查重
	router.GET("/emailOccupy", controller.EmailOccupy)
	// 获取邮箱验证码
	router.POST("/sendEmailCode", controller.SendEmailCode)
	// 注册
	router.POST("/register", controller.Register)
	// 电话号码登录
	router.POST("/phoneLogin", controller.PhoneLogin)
	// 邮箱登录
	router.POST("/emailLogin", controller.EmailLogin)
	// 忘记密码
	router.POST("/forgotPassword", controller.ForgotPassword)
	// 获取用户信息
	router.GET("/userInfo", controller.UserInfo)
	// 退出登录
	router.POST("/logout", controller.Logout)
}
