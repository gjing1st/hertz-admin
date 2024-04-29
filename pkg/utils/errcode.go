// Path: pkg/utils
// FileName: errcode.go
// Created by bestTeam
// Author: GJing
// Date: 2022/11/1$ 20:06$

package utils

// 错误代码100101，其中 10 代表发布平台服务；中间的 01 代表发布平台服务下的文章模块；最后的 01 代表模块下的错误码序号，每个模块可以注册 100 个错误
// 0代表成功
const (
	ErrCode    = 1
	ModuleCode = 100 * ErrCode
	ServerCode = ModuleCode * 100
)

// 服务代码
const (
	ServerPublishCode   = 25 * ServerCode //发布平台
	ServerStatisticCode = 26 * ServerCode //统计服务
)

// 发布平台相关模块
const (
	PublishArticleCode      = 1 * ModuleCode //文章模块
	PublishFileCode         = 2 * ModuleCode //文件模块
	PublishNotificationCode = 3 * ModuleCode //通知模块
)

// 发布平台文章模块错误码
const (
	PublishArticlePublishFailed          = 1 * ErrCode //发布失败
	PublishArticleSaveFailed             = 2 * ErrCode //保存失败
	PublishArticleUploadFailed           = 3 * ErrCode //上传文件失败
	PublishArticleDeleteFailed           = iota + 1    //文章删除失败			4
	PublishArticleQueryFailed                          //查询失败			5
	PublishArticlePushNoticeCreateFailed               //推送通知创建失败		6
	PublishArticleListStandard                         //查询行业规范列表失败	7
	PublishArticleDeleteArticleBatch                   //批量删除文章失败		8
	PublishArticleDeleteArticle                        //删除文章失败			9
	PublishArticleEditArticle                          //编辑文章失败			10
)
