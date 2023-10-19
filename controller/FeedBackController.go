package controller

import (
	"GinCoBlog/request"
	"GinCoBlog/service"
	"GinCoBlog/utils"
	"github.com/gin-gonic/gin"
)

// AddFeedBack 添加用户反馈
func AddFeedBack(c *gin.Context) {
	var param request.AddFeedBackParam
	// 绑定数据
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailResult(c, "参数错误")
		return
	}
	err, tokenInfo := utils.GetCacheUser(c)
	if err != nil {
		utils.AuthorizationResult(c, "登录状态失效")
		return
	}
	// 向数据库内添加
	err, status := service.AddFeedBackService(param, tokenInfo.UserInfo)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 判断是否成功
	if !status {
		utils.FailResult(c, "反馈失败")
		return
	}
	// 反馈成功
	utils.SuccessResult(c, "反馈成功", nil)
}
