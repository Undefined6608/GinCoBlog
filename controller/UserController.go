package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloUser 测试
func HelloUser(c *gin.Context) {
	// 返回数据
	c.JSON(http.StatusOK, gin.H{"msg": "HelloUser!"})
}
