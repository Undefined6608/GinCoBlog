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
	err := c.ShouldBindQuery(params)
	if err != nil {
		utils.FailResult(c, "参数错误")
		return
	}
	// 查询数据库
	if err := service.ArticleListByTypeIdService(&articleList, params.TypeId); err != nil {
		utils.FailResult(c, "获取失败")
		return
	}
	utils.SuccessResult(c, "获取成功", map[string][]entity.Article{"rows": articleList})
}

// ArticleInfoById 文章详情
func ArticleInfoById(c *gin.Context) {
	var params request.ArticleInfoByIdParam
	var info entity.Article
	// 绑定参数
	err := c.ShouldBindQuery(&params)
	if err != nil {
		utils.FailResult(c, "参数错误")
		return
	}
	// 查询数据库
	if err := service.ArticleInfoByIdService(&info, params.ArticleId); err != nil {
		utils.FailResult(c, "获取失败")
		return
	}
	utils.SuccessResult(c, "获取成功", map[string]entity.Article{"data": info})
}

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	// 获取参数实例
	var params request.AddArticleParam
	// 绑定参数
	err := c.ShouldBindJSON(&params)
	if err != nil {
		utils.FailResult(c, "参数错误")
		return
	}
	err, tokenInfo := utils.GetCacheUser(c)
	if err != nil {
		utils.AuthorizationResult(c, "登录状态失效")
		return
	}
	// 将数据添加到数据库
	err, status := service.AddArticleService(params, tokenInfo.UserInfo)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	if !status {
		utils.FailResult(c, "添加失败")
	}
	utils.SuccessResult(c, "添加成功", nil)
}

// UpdateRead 更新阅读量
func UpdateRead(c *gin.Context) {
	var params request.UpdateReadParam
	// 绑定参数
	err := c.ShouldBindJSON(&params)
	if err != nil {
		utils.FailResult(c, "参数错误")
		return
	}
	// 修改数据库
	err, status := service.UpdateReadService(params)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 判断状态
	if !status {
		utils.FailResult(c, "设置失败")
		return
	}
	// 设置成功
	utils.SuccessResult(c, "更新成功", nil)
}

// ArticleComment 获取文章评论
func ArticleComment(c *gin.Context) {
	var params request.ArticleCommentParam
	var commentList []request.ArticleCommentResponse

	// 绑定参数
	err := c.ShouldBindQuery(&params)
	if err != nil {
		utils.FailResult(c, "参数错误")
		return
	}
	// 查询数据库
	if err := service.ArticleCommentService(params, &commentList); err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 获取成功
	utils.SuccessResult(c, "获取成功", map[string][]request.ArticleCommentResponse{"rows": commentList})
}

// AddArticleComment 添加文章评论
func AddArticleComment(c *gin.Context) {
	// 获取参数类型
	var param request.AddArticleCommentParam
	// 绑定参数
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
	// 将数据添加到数据库
	err, status := service.AddArticleCommentService(&param, tokenInfo.UserInfo)
	if err != nil {
		utils.FailResult(c, err.Error())
		return
	}
	// 判断是否添加成功
	if !status {
		utils.FailResult(c, "上传失败")
		return
	}
	utils.SuccessResult(c, "上传成功", nil)
}
