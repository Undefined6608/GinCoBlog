package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// Config 配置 yaml 结构
type Config struct {
	// 端口号
	Port string `yaml:"port"`
	// 数据库配置
	DataBase struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		UserName string `yaml:"user_name"`
		Password string `yaml:"password"`
		Schema   string `yaml:"schema"`
	} `yaml:"data_base"`
	// 日志配置
	LogConfig struct {
		Path string `yaml:"path"`
		Name string `yaml:"name"`
	} `yaml:"log_config"`
	// 正则配置
	Regular struct {
		Phone string `yaml:"phone"`
		Email string `yaml:"email"`
	} `yaml:"regular"`
	// 邮箱配置
	EmailConfig struct {
		EmailAddress string `yaml:"email_address"`
		EmailName    string `yaml:"email_name"`
		Password     string `yaml:"password"`
		SmtpServer   string `yaml:"smtp_server"`
		SmtpPort     int    `yaml:"smtp_port"`
	} `yaml:"email_config"`
	// redis 配置
	RedisConfig struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis_config"`
	// 加密配置
	Encryption struct {
		// 私钥
		PrivateKey struct {
			Password string `yaml:"password"`
		} `yaml:"private_key"`
		// 盐值
		Salt struct {
			Password int `yaml:"password"`
		} `yaml:"salt"`
	} `yaml:"encryption"`
	// token 配置
	Token struct {
		PrivateKey string `yaml:"private_key"`
	} `yaml:"token"`
	// 绕过中间件验证的地址
	NotVerifyUrl []string `yaml:"not_verify_url"`
	// 跨域中间件验证地址
	Cors struct {
		Ip      []string `yaml:"ip"`
		Methods []string `yaml:"methods"`
	} `yaml:"cors"`
	// 上传配置
	Upload struct {
		Host    string `yaml:"host"`
		Path    string `yaml:"path"`
		ImgLoad struct {
			User    string `yaml:"user"`
			Article string `yaml:"article"`
		} `yaml:"img_load"`
		MaxSize struct {
			Img int64 `yaml:"img"`
		} `yaml:"max_size"`
		ImgType []string `yaml:"img_type"`
	} `yaml:"upload"`
}

// Default 获取 yaml 配置
func Default() Config {
	// 实例化配置对象
	var configObj Config
	// 读取配置文件
	yamlFil, err := os.ReadFile("Application.yaml")
	// 读取失败
	if err != nil {
		panic(err)
	}
	// 将读到的文件解析为配置对象
	err = yaml.Unmarshal(yamlFil, &configObj)
	// 解析失败
	if err != nil {
		panic(err)
	}
	// 抛出配置文件对象
	return configObj
}

// Response 返回值类型
type Response struct {
	Code    int         `json:"code"` // 响应值
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 数据
}

// TokenEffectAge Token 生命周期配置
const TokenEffectAge = 1 * 24 * time.Hour
