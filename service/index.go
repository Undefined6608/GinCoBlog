package service

import (
	"GinCoBlog/config"
	"GinCoBlog/entity"
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Pool /** 创建数据库连接池
func Pool() *gorm.DB {
	// 获取数据库配置
	dbConfig := config.Default().DataBase
	// 创建数据库连接
	db, err := gorm.Open(mysql.Open(dbConfig.UserName+":"+dbConfig.Password+"@tcp("+dbConfig.Host+":"+dbConfig.Port+")/"+dbConfig.Schema+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	// 连接失败
	if err != nil {
		panic("连接数据库失败")
	}
	// 抛出数据库连接池
	return db
}

// VerDataBase /** 验证数据库/表是否已经创建
func VerDataBase() {
	// 获取数据库连接池
	pool := Pool()
	// 验证sys_user表是否存在
	err := pool.AutoMigrate(&entity.SysUser{})
	err = pool.AutoMigrate(&entity.ArticleType{})
	err = pool.AutoMigrate(&entity.Article{})
	err = pool.AutoMigrate(&entity.ArticleComments{})
	err = pool.AutoMigrate(&entity.FeedBack{})
	// 表创建失败
	if err != nil {
		panic(err)
	}
}

// RedisClient Redis 数据库连接
func RedisClient() *redis.Client {
	// 拿到 Redis 数据库配置
	redisConfig := config.Default().RedisConfig
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalln(err.Error())
	}
	return rdb
}
