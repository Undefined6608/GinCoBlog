package service

import (
	"GinCoBlog/entity"
	"errors"
)

// ArticleTypeService 文章类型
func ArticleTypeService(list *[]entity.ArticleType) error {
	// 获取数据库内的文章类型数据
	if err := pool.Model(&entity.ArticleType{}).Where("type_visible", 1).Find(&list).Error; err != nil {
		return errors.New("获取失败")
	}
	return nil
}

// ArticleListByTypeId 根据文章类型查询文章列表
func ArticleListByTypeId(list *[]entity.Article, typeId int32) error {
	// 判断参数格式
	if typeId == 0 {
		return errors.New("参数为空")
	}
	// 获取数据库内的文章类型数据
	if err := pool.Model(&entity.Article{}).Where("type_id", typeId).Find(&list).Error; err != nil {
		return errors.New("获取失败")
	}
	return nil
}
