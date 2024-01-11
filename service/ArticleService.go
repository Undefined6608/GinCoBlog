package service

import (
	"GinCoBlog/entity"
	"GinCoBlog/request"
	"GinCoBlog/utils"
	"errors"
	"gorm.io/gorm"
)

// ArticleTypeService 文章类型
func ArticleTypeService(list *[]entity.ArticleType) error {
	// 获取数据库内的文章类型数据
	if err := pool.Model(&entity.ArticleType{}).
		Where("type_visible", 1).
		Find(&list).
		Error; err != nil {
		return errors.New("获取失败")
	}
	return nil
}

// ArticleListByTypeIdService 根据文章类型查询文章列表
func ArticleListByTypeIdService(list *[]request.ArticleListResponse, typeId int32) error {
	// 判断参数格式
	if typeId == 0 {
		return errors.New("参数为空")
	}
	// 获取数据库内的文章类型数据
	if err := pool.
		Table("article").
		Select("article.id, article.type_id, article.user_id, article.title, article.context, article.article_visible, "+
			"article.read, article.create_time as date, article.icon, sys_user.user_name, sys_user.head_sculpture, "+
			"sys_user.integral, sys_user.member, sys_user.uuid").
		Joins("JOIN sys_user ON article.user_id = sys_user.uid").
		Where("article.type_id = ? AND article.article_visible = ?", typeId, 1).
		Find(&list).
		Error; err != nil {
		return errors.New("获取失败")
	}
	return nil
}

// ArticleInfoByIdService 根据 ID 获取文章详情
func ArticleInfoByIdService(info *entity.Article, articleId int32) error {
	if articleId == 0 {
		return errors.New("参数为空")
	}
	// 获取数据库信息
	if err := pool.
		Model(&entity.Article{}).
		Where("id=? AND article_visible=?", articleId, 1).
		First(&info).
		Error; err != nil {
		return errors.New("获取失败")
	}
	return nil
}

// AddArticleService 添加文章
func AddArticleService(param request.AddArticleParam, userInfo entity.SysUser) (error, bool) {
	// 判断参数是否为空
	if param.TypeId == 0 || utils.StrIsEmpty(param.Title) || utils.StrIsEmpty(param.Context) || utils.StrIsEmpty(param.Icon) {
		return errors.New("参数为空"), false
	}
	// 判断参数是否合法
	// 将数据存入数据库
	err := pool.Model(&entity.Article{}).Create(&entity.Article{
		TypeId:  param.TypeId,
		UserId:  userInfo.UID,
		Title:   param.Title,
		Context: param.Context,
		Icon:    param.Icon,
	}).Error
	if err != nil {
		return errors.New("添加失败"), false
	}
	return nil, true
}

// UpdateReadService 更新阅读量
func UpdateReadService(param request.UpdateReadParam) (error, bool) {
	// 判断参数为空
	if param.ArticleId == 0 {
		return errors.New("参数为空"), false
	}
	// 修改数据库
	if err := pool.Model(&entity.Article{}).Where("id", param.ArticleId).UpdateColumn("read", gorm.Expr("`read` + ?", 1)).Error; err != nil {
		return errors.New("修改失败"), false
	}
	// 修改成功
	return nil, true
}

// ArticleCommentService 获取文章评论
func ArticleCommentService(param request.ArticleCommentParam, comments *[]request.ArticleCommentResponse) error {
	if param.ArticleId == 0 {
		return errors.New("参数为空")
	}
	// 获取数据库信息
	if err := pool.Table("article_comments").
		Where("article_id", param.ArticleId).
		Joins("INNER JOIN sys_user ON article_comments.user_id = sys_user.uid").
		Select("article_comments.id, article_comments.article_id, article_comments.context, sys_user.user_name, sys_user.head_sculpture, sys_user.integral, sys_user.member").
		Find(&comments).Error; err != nil {
		return errors.New("获取失败")
	}
	// 获取成功
	return nil
}

// AddArticleCommentService 添加文章评论
func AddArticleCommentService(param *request.AddArticleCommentParam, userInfo entity.SysUser) (error, bool) {
	// 判断参数是否为空
	if param.ArticleId == 0 || utils.StrIsEmpty(param.Context) {
		return errors.New("参数为空"), false
	}
	// 判断参数格式
	// 将数据存入数据库
	if err := pool.Model(&entity.ArticleComments{}).Create(&entity.ArticleComments{
		ArticleID: param.ArticleId,
		UserID:    userInfo.UID,
		Context:   param.Context,
	}).Error; err != nil {
		return errors.New("上传失败"), false
	}
	// 添加成功
	return nil, true
}
