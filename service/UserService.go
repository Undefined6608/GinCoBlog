package service

import (
	"GinCoBlog/config"
	"GinCoBlog/entity"
	"GinCoBlog/request"
	"GinCoBlog/utils"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

// 获取数据库连接池
var pool = Pool()

// 定义上下文对象
var (
	ctx = context.Background()
)

// UserNameOccupyService 用户名查重
func UserNameOccupyService(userName string) (error, bool) {
	// 创建数据库数组对象
	var user []entity.SysUser
	// 验证用户名参数
	// 验证是否为空
	if utils.StrIsEmpty(userName) {
		return errors.New("参数为空"), false
	}
	// 判断长度是否超出
	if len(userName) > 15 {
		return errors.New("参数超出 15 个字符"), false
	}
	// 判断用户名中是否含有非法字符串
	// 通过调用 gorm api 进行查找数据库
	pool.Where("user_name=?", userName).Find(&user)
	// 判断查询的数组长度是否大于 1
	if len(user) > 1 {
		// 将错误写入日志
		logrus.Error("数据库内存有多个同名用户名")
		// 将信息返回
		return errors.New("用户名已存在"), false
	}
	// 返回用户是否存在
	return nil, len(user) == 1
}

// PhoneOccupyService 电话号码查重
func PhoneOccupyService(phone string) (error, bool) {
	// 创建数据库数组对象
	var user []entity.SysUser
	// 验证用户名参数
	// 验证是否为空
	if utils.StrIsEmpty(phone) {
		return errors.New("参数为空"), false
	}
	// 判断电话号码格式
	if utils.VerPhoneReg(phone) {
		return errors.New("电话号码格式错误"), false
	}
	// 判断用户名中是否含有非法字符串
	// 通过调用 gorm api 进行查找数据库
	pool.Where("phone=?", phone).Find(&user)
	// 判断查询的数组长度是否大于 1
	if len(user) > 1 {
		// 将错误写入日志
		logrus.Error("数据库内存有多个相同的电话号码")
		// 将信息返回
		return errors.New("电话号码已存在"), false
	}
	// 返回用户是否存在
	return nil, len(user) == 1
}

// EmailOccupyService 邮箱查重
func EmailOccupyService(email string) (error, bool) {
	// 创建数据库数组对象
	var user []entity.SysUser
	// 验证用户名参数
	// 验证是否为空
	if utils.StrIsEmpty(email) {
		return errors.New("参数为空"), false
	}
	// 判断邮箱格式是否正确
	if utils.VerEmailReg(email) {
		return errors.New("邮箱格式错误"), false
	}
	// 判断用户名中是否含有非法字符串
	// 通过调用 gorm api 进行查找数据库
	pool.Where("email=?", email).Find(&user)
	// 判断查询的数组长度是否大于 1
	if len(user) > 1 {
		// 将错误写入日志
		logrus.Error("数据库内存有多个相同的邮箱")
		// 将信息返回
		return errors.New("邮箱已存在"), false
	}
	// 返回用户是否存在
	return nil, len(user) == 1
}

// SendMsgCodeService 发送邮箱验证码
func SendMsgCodeService(email string) (error, bool) {
	// 验证用户名参数
	// 验证是否为空
	if utils.StrIsEmpty(email) {
		return errors.New("参数为空"), false
	}
	// 判断邮箱格式是否正确
	if utils.VerEmailReg(email) {
		return errors.New("邮箱格式错误"), false
	}
	// 判断用户名中是否含有非法字符串
	// 发送验证码
	code := utils.SendEmail(email)
	// 判断生成的验证码格式是否正确
	if len(code) != 6 {
		return errors.New("验证码发送失败"), false
	}
	// 将 code 存入 redis 5 分钟有效期
	if _, err := RedisClient().Set(ctx, email, code, 5*time.Minute).Result(); err != nil {
		return errors.New("发送失败"), false
	}
	// 返回
	return nil, true
}

// RegisterService 注册
func RegisterService(param *request.RegisterParams) (error, bool) {
	// 判断数据是否为空
	if utils.StrIsEmpty(param.UserName) || utils.StrIsEmpty(param.Phone) || utils.StrIsEmpty(param.Email) || utils.StrIsEmpty(param.Password) {
		return errors.New("参数为空"), false
	}
	// 验证电话号码格式
	if utils.VerPhoneReg(param.Phone) {
		return errors.New("电话号码格式错误"), false
	}
	// 验证邮箱格式
	if utils.VerEmailReg(param.Email) {
		return errors.New("邮箱格式错误"), false
	}
	// 验证电话号码是否已存在
	if err, _ := PhoneOccupyService(param.Phone); err != nil {
		return err, false
	}
	// 验证邮箱是否已存在
	if err, _ := EmailOccupyService(param.Email); err != nil {
		return err, false
	}
	// 判断密码格式
	if len(param.Password) != 32 || len(param.VerPassword) != 32 {
		return errors.New("密码为空"), false
	}
	// 判断两次密码是否一致
	if param.Password != param.VerPassword {
		return errors.New("两次密码不一致"), false
	}
	// 判断验证码是否正确
	result, err := RedisClient().Get(ctx, param.Email).Result()
	if err != nil {
		return errors.New("验证码失效"), false
	}
	if result != param.EmailCode {
		return errors.New("验证码错误！"), false
	}
	// 将 redis 中验证码删除
	if _, err := RedisClient().Del(ctx, param.Email).Result(); err != nil {
		return errors.New("注册失败"), false
	}
	// 将密码进行二次加密
	encryptionPwd := utils.EncryptionPassword(param.Password)
	// 将用户数据存入数据库
	if err := pool.Create(&entity.SysUser{
		UserName: param.UserName,
		Phone:    param.Phone,
		Password: encryptionPwd,
		Email:    param.Email,
		UUID:     utils.CreateUUID(),
	}).Error; err != nil {
		return errors.New("注册失败！"), false
	}
	// 判断
	return nil, true
}

// PhoneLoginService 电话号码登录
func PhoneLoginService(param *request.PhoneLoginParams) (error, bool, string) {
	// 创建数据库数组对象
	var user []entity.SysUser
	// 验证参数为空
	if utils.StrIsEmpty(param.Phone) || utils.StrIsEmpty(param.Password) {
		return errors.New("参数为空"), false, ""
	}
	// 验证电话号码格式
	if utils.VerPhoneReg(param.Phone) {
		return errors.New("电话号码格式错误"), false, ""
	}
	// 验证密码格式
	if len(param.Password) != 32 {
		return errors.New("密码格式错误"), false, ""
	}
	// 判断此用户是否存在
	// 通过电话号码查询用户
	err := pool.Where("phone=?", param.Phone).Find(&user).Error
	if err != nil {
		return errors.New("登录失败"), false, ""
	}
	// 通过判断用户列表长度，判断是否有此用户
	if len(user) < 1 {
		return errors.New("此用户暂未注册！"), false, ""
	}
	// 验证密码正确性
	if !utils.ComparePassword(user[0].Password, param.Password) {
		return errors.New("账号/密码错误！"), false, ""
	}
	if user[0].Available {
		return errors.New("账号不存在/异常"), false, ""
	}
	// 生成 Token
	token := utils.GenerateToken(&request.TokenParams{
		UserInfo:       user[0],
		StandardClaims: jwt.StandardClaims{},
	})
	if utils.StrIsEmpty(token) {
		return errors.New("登录失败"), false, ""
	}
	log.Println("user")
	// 将 Token 存入 Redis
	if _, err := RedisClient().Set(ctx, token, param.Phone, config.TokenEffectAge).Result(); err != nil {
		return errors.New("登录失败"), false, ""
	}

	return nil, true, token
}

// EmailLoginService 邮箱登录
func EmailLoginService(param *request.EmailLoginParams) (error, bool, string) {
	// 创建数据库数组对象
	var user []entity.SysUser
	// 验证参数为空
	if utils.StrIsEmpty(param.Email) || utils.StrIsEmpty(param.Password) {
		return errors.New("参数为空"), false, ""
	}
	// 验证邮箱格式
	if utils.VerEmailReg(param.Email) {
		return errors.New("邮箱格式错误"), false, ""
	}
	// 验证密码格式
	if len(param.Password) != 32 {
		return errors.New("密码格式错误"), false, ""
	}
	// 判断此用户是否存在
	// 通过邮箱查询用户
	err := pool.Where("email=?", param.Email).Find(&user).Error
	if err != nil {
		return errors.New("登录失败"), false, ""
	}
	// 通过判断用户列表长度，判断是否有此用户
	if len(user) < 1 {
		return errors.New("此用户暂未注册！"), false, ""
	}
	// 验证密码正确性
	if !utils.ComparePassword(user[0].Password, param.Password) {
		return errors.New("账号/密码错误！"), false, ""
	}
	if user[0].Available {
		return errors.New("账号不存在/异常"), false, ""
	}
	// 生成 Token
	token := utils.GenerateToken(&request.TokenParams{
		UserInfo:       user[0],
		StandardClaims: jwt.StandardClaims{},
	})
	if utils.StrIsEmpty(token) {
		return errors.New("登录失败"), false, ""
	}
	// 将 Token 存入 Redis
	if _, err := RedisClient().Set(ctx, token, param.Email, config.TokenEffectAge).Result(); err != nil {
		return errors.New("登录失败"), false, ""
	}

	return nil, true, token
}

// ForgotPasswordService ForgotPassword 忘记密码
func ForgotPasswordService(param *request.ForgotPasswordParams) (error, bool) {
	// 创建数据库数组对象
	var user []entity.SysUser
	// 验证参数为空
	if utils.StrIsEmpty(param.Email) || utils.StrIsEmpty(param.EmailCode) || utils.StrIsEmpty(param.NewPassword) || utils.StrIsEmpty(param.VerPassword) {
		return errors.New("参数为空"), false
	}
	// 验证邮箱格式
	if utils.VerEmailReg(param.Email) {
		return errors.New("邮箱格式错误"), false
	}
	// 验证密码格式
	if len(param.NewPassword) != 32 || len(param.VerPassword) != 32 {
		return errors.New("密码格式错误"), false
	}
	// 验证新密码和验证密码是否一致
	if param.NewPassword != param.VerPassword {
		return errors.New("两次密码不一致"), false
	}
	// 验证验证码格式
	if len(param.EmailCode) != 6 {
		return errors.New("验证码格式错误"), false
	}
	// 验证验证码是否正确
	if result, err := RedisClient().Get(ctx, param.Email).Result(); err != nil || result != param.EmailCode {
		return errors.New("验证码错误"), false
	}
	// 删除 redis 中的验证码
	if _, err := RedisClient().Del(ctx, param.Email).Result(); err != nil {
		return errors.New("修改失败"), false
	}
	// 判断此用户是否存在
	// 通过邮箱查询用户
	err := pool.Where("email=?", param.Email).Find(&user).Error
	if err != nil {
		return errors.New("修改失败"), false
	}
	// 通过判断用户列表长度，判断是否有此用户
	if len(user) < 1 {
		return errors.New("账号不存在"), false
	}
	// 对密码进行加密
	encryptionPassword := utils.EncryptionPassword(param.NewPassword)
	// 修改密码 SQL
	err = pool.Model(&entity.SysUser{}).Where("uid", user[0].UID).Update("password", encryptionPassword).Error
	if err != nil {
		return errors.New("修改失败"), false
	}
	// 修改成功
	return nil, true
}

// VerUserByToken 通过Token查找用户
func VerUserByToken(token string) (string, error) {
	// 验证数据库中是否存有此token
	user, err := RedisClient().Get(ctx, token).Result()
	return user, err
}

// LogoutService 退出登录
func LogoutService(param *request.TokenParams, token string) (error, bool) {
	// 判断参数是否为空
	if param == nil {
		return errors.New("退出失败"), false
	}
	// 删除 Redis
	_, err := RedisClient().Del(ctx, token).Result()
	if err != nil {
		return err, false
	}
	return nil, true
}
