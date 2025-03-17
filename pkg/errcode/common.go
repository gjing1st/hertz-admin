// Path: pkg/errcode
// FileName: common.go
// Created by bestTeam
// Author: GJing
// Date: 2023/6/6$ 16:04$

package errcode

// CommonCode 此处为通用错误代码4位
const CommonCode = ErrCode * 1000

const (
	OperationFail = CommonCode + 100 + iota //操作失败
	ValueErr                                //参数错误
)

// 数据库类错误
const (
	DBErr       = CommonCode + iota //数据库操作错误
	DBCreateErr                     //数据库操作错误
	DBFindErr                       //数据库查询错误
	DBUpdateErr
	DBDeleteErr
)

// 用户错误
const (
	UserNotFound  = CommonCode*2 + iota //用户不存在
	PasswordErr                         //用户名/密码错误
	MAXAdmin                            //已达最大用户数
	UserNameExist                       //用户已存在
	NoToken                             //未携带token
	TokenExpired                        //token过期
	RoleErr                             //权限错误
	ForbiddenErr                        //无权访问
	PwdErr                              //密码错误
	PwdMaxErr                           //密码错误次数过多已锁定
)

// 系统运行中的错误
const (
	SysCacheErr         = 3*CommonCode + iota //缓存错误
	SysCacheGetErr                            //缓存获取错误
	SysCacheSetErr                            //缓存获取错误
	SysCmdErr                                 //执行宿主机cmd指令出错
	SysNetworkErr                             //配置网卡错误
	SysJsonMarshalErr                         //转json失败
	SysJsonUnMarshalErr                       //json解析失败
	SysTimeParseErr                           //时间转换错误
	SysSaveFileErr                            //保存文件错误
)

const (
	IdCardErr       = 4*CommonCode + iota
	IdCardNumberErr //身份证号错误
)
