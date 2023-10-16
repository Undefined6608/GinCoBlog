package controller

import (
	"GinCoBlog/entity"
	"GinCoBlog/request"
	"GinCoBlog/service"
	"GinCoBlog/utils"
	"github.com/gin-gonic/gin"
	"sort"
)

// ArticleType 文章类型
func ArticleType(c *gin.Context) {
	// 定义类型
	var typeList []entity.ArticleType
	// 获取数据
	if err := service.ArticleTypeService(&typeList); err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 对数据进行排序
	sort.Slice(typeList, func(i, j int) bool {
		return typeList[i].Order < typeList[j].Order
	})
	// 返回数据
	utils.SuccessResult(c, "获取成功", map[string][]entity.ArticleType{"articleType": typeList})
}

// ArticleListByType 文章列表
func ArticleListByType(c *gin.Context) {
	var params request.ArticleListByTypeParam
	var articleList []entity.Article
	// 绑定参数
	err := c.ShouldBindJSON(&params)
	if err != nil {
		utils.FailResult(c, "参数错误")
		return
	}
	// 查询数据库
	if err := service.ArticleListByTypeId(&articleList, params.TypeId); err != nil {
		utils.FailResult(c, "获取失败")
		return
	}
	utils.SuccessResult(c, "获取成功", map[string][]entity.Article{"rows": articleList})
}
