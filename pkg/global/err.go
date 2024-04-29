package global

import "errors"

var (
	RequestParamErr   = errors.New("invalid request param")
	ServerErr         = errors.New("server error")
	RequestErrExt     = errors.New("上传包格式错误")
	UsernameHasExists = errors.New("该用户名已存在")
	TokenExpired      = errors.New("token is expired or error")
	UserNotExisted    = errors.New("user not existed")
	AdminPasswordErr  = errors.New("password error")
	AdminDisabledErr  = errors.New("admin account disabled")
)

var (
	DBNullErr   = errors.New("缺少数据库实例")
	InitDataErr = errors.New("初始化数据失败")
)
