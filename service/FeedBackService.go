package service

import (
	"GinCoBlog/entity"
	"GinCoBlog/request"
	"GinCoBlog/utils"
	"errors"
)

// AddFeedBackService 添加用户反馈
func AddFeedBackService(param request.AddFeedBackParam, userInfo entity.SysUser) (error, bool) {
	// 验证数据
	if utils.StrIsEmpty(param.Context) {
		return errors.New("参数为空"), false
	}
	// 验证数据
	// 向数据库内添加数据
	if err := pool.Model(&entity.FeedBack{}).Create(&entity.FeedBack{
		UserID:  userInfo.UID,
		Context: param.Context,
	}).Error; err != nil {
		return errors.New("上传失败"), false
	}
	return nil, true
}
