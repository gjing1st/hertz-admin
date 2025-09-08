// Path: pkg/errcode
// FileName: string.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/28$ 21:44$

package errcode

import "errors"

func (e Err) Error() string {
	switch {
	//系统运行相关
	case errors.Is(e, SuccessCode):
		return "操作成功"
	case errors.Is(e, HaSysErr):
		return "系统错误"
	case errors.Is(e, HaSysJsonMarshalErr):
		return "转换为Json格式错误"
	case errors.Is(e, HaSysJsonUnMarshalErr):
		return "json解析错误"
	case errors.Is(e, HaSysTimeParseErr):
		return "时间解析错误"
	case errors.Is(e, HaSysParamErr):
		return "参数错误"
		//用户相关
	case errors.Is(e, HaUserNotLogin):
		return "用户未登录或登录已失效，请重新登录"
		//用户相关
	case errors.Is(e, NoToken), errors.Is(e, TokenExpired):
		return "请重新登录"
		// 系统运行中的缓存错误
	case errors.Is(e, SysCacheErr):
		return "缓存错误"
	case errors.Is(e, SysCacheGetErr):
		return "缓存获取错误"
	case errors.Is(e, SysCacheSetErr):
		return "缓存设置错误"
		// 系统运行中数据库出现的错误
	case errors.Is(e, SysCmdErr):
		return "执行宿主机cmd指令出错"
	case errors.Is(e, SysNetworkErr):
		return "网卡配置出错"
		// 系统运行中的错误
	case errors.Is(e, SysJsonMarshalErr):
		return "转json失败"
	case errors.Is(e, SysJsonUnMarshalErr):
		return "json解析失败"
	case errors.Is(e, SysTimeParseErr):
		return "时间转换错误"
	case errors.Is(e, SysSaveFileErr):
		return "保存文件错误"
	case errors.Is(e, UserNotFound):
		return "用户不存在"
	case errors.Is(e, PasswordErr):
		return "用户名或密码错误"
	case errors.Is(e, PwdErr):
		return "密码输入错误，剩余次数："
	case errors.Is(e, PwdMaxErr):
		return "密码错误次数过多，账号已锁定"
	case errors.Is(e, MAXAdmin):
		return "管理员数量已达最大限制"
	case errors.Is(e, UserNameExist):
		return "当前用户已注册，新增失败"
	case errors.Is(e, DBErr), errors.Is(e, DBCreateErr), errors.Is(e, DBFindErr), errors.Is(e, DBUpdateErr), errors.Is(e, DBDeleteErr):
		return "db-系统错误"
	case errors.Is(e, OperationFail):
		return "操作失败"
	case errors.Is(e, ValueErr):
		return "数值错误"
	case errors.Is(e, IdCardNumberErr):
		return "身份证号错误"
	}

	return "未知错误"
}
