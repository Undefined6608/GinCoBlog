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
	if err := pool.Model(&entity.ArticleType{}).Where("type_visible", 1).Find(&list).Error; err != nil {
		return errors.New("获取失败")
	}
	return nil
}

// ArticleListByTypeIdService 根据文章类型查询文章列表
func ArticleListByTypeIdService(list *[]entity.Article, typeId int32) error {
	// 判断参数格式
	if typeId == 0 {
		return errors.New("参数为空")
	}
	// 获取数据库内的文章类型数据
	if err := pool.Model(&entity.Article{}).Select("id", "type_id", "user_id", "title", "read", "create_time", "icon").Where("type_id=? AND article_visible=?", typeId, 1).Find(&list).Error; err != nil {
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
	if err := pool.Model(&entity.Article{}).Where("id=? AND article_visible=?", articleId, 1).First(&info).Error; err != nil {
		return errors.New("获取失败")
	}
	return nil
}

// AddArticleService 添加文章
func AddArticleService(param *request.AddArticleParam, userInfo entity.SysUser) (error, bool) {
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
func UpdateReadService(param *request.UpdateReadParam) (error, bool) {
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
