package middleware

import (
	"GinCoBlog/config"
	"GinCoBlog/request"
	"GinCoBlog/service"
	"GinCoBlog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
)

// JwtVerifyMiddle /**验证token
func JwtVerifyMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		//过滤是否验证token
		if utils.IsContainArr(config.NotVerifyUrl, c.Request.RequestURI) {
			return
		}
		token := c.GetHeader("Authorization")
		if token == "" {
			log.Println("")
			utils.AuthorizationResult(c, "权限验证失败,无法访问系统资源！")
			// 终止请求
			c.Abort()
			return
		}
		// 判断是否错误
		claims := parseToken(c, token)
		//验证token，并存储在请求中
		c.Set("user", claims)
	}
}

// 解析Token
func parseToken(c *gin.Context, tokenString string) *request.TokenParams {
	// 验证数据库中是否存有此token
	user, err := service.VerUserByToken(tokenString)
	// 没查到
	if err != nil || user == "" {
		utils.AuthorizationResult(c, "权限验证失败,无法访问系统资源！")
		c.Abort()
		return nil
	}
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &request.TokenParams{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenPrivateKey), nil
	})
	// 解析成功
	if err != nil {
		utils.AuthorizationResult(c, "权限验证失败,无法访问系统资源！")
		c.Abort()
		return nil
	}
	// 将Token内存的数据转化为 token.Claims
	claims, ok := token.Claims.(*request.TokenParams)
	// 转化失败
	if !ok {
		utils.AuthorizationResult(c, "登录失效！")
		c.Abort()
		return nil
	}
	// 抛出数据
	return claims
}
