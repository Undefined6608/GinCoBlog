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

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "HelloWorld!")
	})

	err := router.Run(":4001")
	if err != nil {
		panic("项目启动失败!")
		return
	}
}
