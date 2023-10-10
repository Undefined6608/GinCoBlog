package controller

import (
	"GinCoBlog/request"
	"GinCoBlog/service"
	"GinCoBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloUser 测试
func HelloUser(c *gin.Context) {
	// 返回数据
	c.JSON(http.StatusOK, gin.H{"msg": "HelloUser!"})
}

// UserNameOccupy 用户名查重
func UserNameOccupy(c *gin.Context) {
	// 获取参数接口实例
	var param request.UserNameOccupyParams
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 参数绑定失败
	if err != nil {
		utils.FailResult(c, "参数错误！")
		return
	}
	// 调用接口查找用户名是否已存在
	err, isUserName := service.UserNameOccupyService(param.UserName)
	// 如果发生错误
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 判断用户名是否存在
	if isUserName {
		utils.FailResult(c, "用户名已存在！")
		return
	}
	// 如果用户名不存在
	utils.SuccessResult(c, "可以使用", nil)
}

// PhoneOccupy 电话号码查重
func PhoneOccupy(c *gin.Context) {
	// 获取参数接口实例
	var param request.PhoneOccupyParams
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 参数绑定失败
	if err != nil {
		utils.FailResult(c, "参数错误！")
		return
	}
	// 调用接口查找用户名是否已存在
	err, isUserName := service.PhoneOccupyService(param.Phone)
	// 如果发生错误
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 判断用户名是否存在
	if isUserName {
		utils.FailResult(c, "电话号码已存在！")
		return
	}
	// 如果用户名不存在
	utils.SuccessResult(c, "可以使用", nil)
}

// EmailOccupy 邮箱查重
func EmailOccupy(c *gin.Context) {
	// 获取参数接口实例
	var param request.EmailOccupyParams
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 参数绑定失败
	if err != nil {
		utils.FailResult(c, "参数错误！")
		return
	}
	// 调用接口查找用户名是否已存在
	err, isUserName := service.EmailOccupyService(param.Email)
	// 如果发生错误
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 判断用户名是否存在
	if isUserName {
		utils.FailResult(c, "邮箱已存在！")
		return
	}
	// 如果用户名不存在
	utils.SuccessResult(c, "可以使用", nil)
}

// SendEmailCode 获取邮箱验证码
func SendEmailCode(c *gin.Context) {
	// 获取参数接口实例
	var param request.SendEmailParams
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 参数绑定失败
	if err != nil {
		utils.FailResult(c, "参数错误！")
		return
	}
	// 调用接口发送验证码
	err, sendStatus := service.SendMsgCodeService(param.Email)
	// 如果发生错误
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 发送失败
	if !sendStatus {
		utils.FailResult(c, "发送失败！")
		return
	}
	// 发送成功
	utils.SuccessResult(c, "发送成功！", nil)
}

func Register(c *gin.Context) {
	// 获取参数接口实例
	var param request.RegisterParams
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 参数绑定失败
	if err != nil {
		utils.FailResult(c, "参数错误！")
		return
	}
	// 验证完成后
	err, status := service.RegisterService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 注册失败
	if !status {
		utils.FailResult(c, "注册失败")
		return
	}
	// 注册成功
	utils.SuccessResult(c, "注册成功！", nil)
}
