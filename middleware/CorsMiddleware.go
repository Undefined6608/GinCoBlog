package middleware

import (
	"GinCoBlog/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsMiddleware 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	allowAccess := config.Cors.Ip
	corsConfig.AllowOrigins = allowAccess // 允许访问的域名
	corsConfig.AllowMethods = config.Cors.Methods

	return cors.New(corsConfig)
}
