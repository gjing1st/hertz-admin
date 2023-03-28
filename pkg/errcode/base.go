// Path: pkg/errcode
// FileName: err.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/28$ 21:11$

package errcode

type Err int32

// 错误代码100101，其中 10 代表发布平台服务；中间的 01 代表发布平台服务下的文章模块；最后的 01 代表模块下的错误码序号，每个模块可以注册 100 个错误
// 0代表成功
const (
	SuccessCode Err = 0
	ErrCode     Err = 1
	ModuleCode      = 100 * ErrCode
	ServerCode      = ModuleCode * 100
)

// 模块代码
const (
	HaSysCode  = 1 * ModuleCode //系统中产生的错误代码
	HaInitCode = 2 * ModuleCode //初始化模块
	HaUserCode = 3 * ModuleCode //用户信息模块
	HaLogCode  = 6 * ModuleCode //日志模块
)

// 服务平台
const (
	HaServer   = 10 * ServerCode // Ha管理平台
	AuthServer = 11 * ServerCode // 认证服务
)

// 系统运行中的错误
const (
	HaSysErr              = 1<<iota + HaServer + HaSysCode
	HaSysJsonMarshalErr   //转json失败
	HaSysJsonUnMarshalErr //json解析失败
	HaSysTimeParseErr     //时间转换错误
)

const (
	HaUserNotLogin = 1<<iota + HaServer + HaUserCode
)
