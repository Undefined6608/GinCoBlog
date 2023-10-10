package service

import (
	"GinCoBlog/entity"
	"GinCoBlog/request"
	"GinCoBlog/utils"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
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
	// 判断密码格式
	if len(param.Password) != 32 || len(param.VerPassword) != 32 {
		return errors.New("密码为空"), false
	}
	// 判断两次密码是否一致
	if param.Password != param.VerPassword {
		return errors.New("两次密码不一致"), false
	}
	// 判断验证码是否正确
	/*result, err := RedisClient().Get(ctx, param.Email).Result()
	if err != nil {
		return err, false
	}*/

	// 判断
	return nil, true
}
