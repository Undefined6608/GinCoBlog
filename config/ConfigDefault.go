package config

import (
	"gopkg.in/yaml.v2"
	"os"
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
		Path string `yaml:"path"` // 日志存放路径
		Name string `yaml:"name"` // 日志名称
	} `yaml:"log_config"`
}

// Default /** 获取 yaml 配置
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

// Response /**返回值类型
type Response struct {
	Code    int         `json:"code"` // 响应值
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 数据
}

// LogFilePath /**日志存放地址
var LogFilePath = Default().LogConfig.Path

// LogFileName /**日志文件名
var LogFileName = Default().LogConfig.Name
