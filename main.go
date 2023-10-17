/**
 * @projectName:    GinCoBlog
 * @package:        GinCoBlog
 * @className:      main
 * @author:     张杰
 * @description:  TODO
 * @date:    2023/10/4 9:23
 * @version:    1.0
 */
package main

// 引入相关依赖
import (
	"GinCoBlog/config"
	"GinCoBlog/middleware"
	routes "GinCoBlog/router"
	"GinCoBlog/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 项目主方法
func main() {
	// 获取项目路由
	router := gin.New()
	// 挂载中间件
	router.Use(gin.Logger(), gin.Recovery(), middleware.CorsMiddleware(), middleware.LoggerToFile(), middleware.JwtVerifyMiddle())
	// 加载代理中间件
	err := router.SetTrustedProxies([]string{"192.168.1.0/24"})
	if err != nil {
		fmt.Println("代理失败！")
		return
	}
	// 验证数据表是否存在
	service.VerDataBase()
	// 编写项目基础接口
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "HelloWorld!")
	})
	// 调用项目主路由
	routes.SetupRouterGroup(router.Group("/api"))
	// 开启端口监听
	err = router.Run(":" + config.Default().Port)
	// 开启监听失败
	if err != nil {
		// 写入日志
		log.Fatalln("项目启动失败!")
		return
	}
}
