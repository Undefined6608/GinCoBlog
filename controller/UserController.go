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

// Register 注册
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

// PhoneLogin 电话号码登录
func PhoneLogin(c *gin.Context) {
	// 拿到电话号码登录参数实体类
	var param request.PhoneLoginParams
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 参数绑定失败
	if err != nil {
		utils.FailResult(c, "参数错误！")
		return
	}
	// 验证完成后
	err, status, token := service.PhoneLoginService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 注册失败
	if !status {
		utils.FailResult(c, "登录失败")
		return
	}
	// 注册成功
	utils.SuccessResult(c, "登录成功！", map[string]interface{}{"token": token})
}

// EmailLogin 邮箱登录
func EmailLogin(c *gin.Context) {
	// 拿到电话号码登录参数实体类
	var param request.EmailLoginParams
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 参数绑定失败
	if err != nil {
		utils.FailResult(c, "参数错误！")
		return
	}
	// 验证完成后
	err, status, token := service.EmailLoginService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 注册失败
	if !status {
		utils.FailResult(c, "登录失败")
		return
	}
	// 注册成功
	utils.SuccessResult(c, "登录成功！", map[string]interface{}{"token": token})
}

// ForgotPassword 忘记密码
func ForgotPassword(c *gin.Context) {
	// 获取参数接口实例
	var param request.ForgotPasswordParams
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 参数绑定失败
	if err != nil {
		utils.FailResult(c, "参数错误！")
		return
	}
	// 验证完成后
	err, status := service.ForgotPasswordService(&param)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 注册失败
	if !status {
		utils.FailResult(c, "修改失败")
		return
	}
	// 注册成功
	utils.SuccessResult(c, "修改成功！", nil)
}

// UserInfo 用户信息
func UserInfo(c *gin.Context) {
	err, tokenParam := utils.GetCacheUser(c)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 返回用户信息
	utils.SuccessResult(c, "获取成功！", request.UserResponse{
		Name:          tokenParam.UserInfo.UserName,
		Phone:         tokenParam.UserInfo.Phone,
		HeadSculpture: tokenParam.UserInfo.HeadSculpture,
		Email:         tokenParam.UserInfo.Email,
		Limit:         tokenParam.UserInfo.Limit,
		Integral:      tokenParam.UserInfo.Integral,
		Member:        tokenParam.UserInfo.Member,
		CreateTime:    tokenParam.UserInfo.CreateTime.String(),
		UpdateTime:    tokenParam.UserInfo.UpdateTime.String(),
		UUID:          tokenParam.UserInfo.UUID,
	})
}

// ModifyUserInfo 修改用户信息
func ModifyUserInfo(c *gin.Context) {
	var params request.ModifyUserInfoParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		utils.FailResult(c, "参数错误")
		return
	}
	err, tokenInfo := utils.GetCacheUser(c)
	if err != nil {
		utils.AuthorizationResult(c, err.Error())
		return
	}
	err, status := service.ModifyUserInfoService(&params, &tokenInfo.UserInfo)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "修改失败")
		return
	}
	utils.SuccessResult(c, "修改成功", nil)
}

// ModifyPassword 修改密码
func ModifyPassword(c *gin.Context) {
	// 获取参数类型对象
	var params request.ModifyPasswordParams
	// 绑定参数
	err := c.ShouldBindJSON(&params)
	// 绑定失败
	if err != nil {
		utils.FailResult(c, "参数错误")
		return
	}
	// 获取Token信息
	err, tokenInfo := utils.GetCacheUser(c)
	if err != nil {
		utils.AuthorizationResult(c, "登录失效")
		return
	}
	err, status := service.ModifyPasswordService(&params, &tokenInfo.UserInfo)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "修改失败")
		return
	}
	utils.SuccessResult(c, "修改成功", nil)
}

// Logout 退出登录
func Logout(c *gin.Context) {
	// 获取用户信息
	user, _ := c.Get("user")
	// 获取Token
	header := c.GetHeader("Authorization")
	// 判断用户信息是否存在
	if user == nil {
		utils.FailResult(c, "登陆失效，请重新登录！")
		return
	}
	// 将user转化为 TokenParam类型
	tokenParam, ok := user.(*request.TokenParams)
	// 判断是否转化正确
	if !ok {
		utils.FailResult(c, "登陆失效，请重新登录！")
		return
	}
	err, logout := service.LogoutService(tokenParam, header)
	if err != nil {
		utils.FailResult(c, "退出失败")
		return
	}
	if !logout {
		utils.FailResult(c, "退出失败")
		return
	}
	// 退出成功
	utils.SuccessResult(c, "退出成功", nil)
}
