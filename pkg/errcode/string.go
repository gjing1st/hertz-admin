// Path: pkg/errcode
// FileName: string.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/28$ 21:44$

package errcode

func (e Err) String() string {
	switch e {
	//系统运行相关
	case SuccessCode:
		return "操作成功"
	case HaSysErr:
		return "系统错误"
	case HaSysJsonMarshalErr:
		return "转换为Json格式错误"
	case HaSysJsonUnMarshalErr:
		return "json解析错误"
	case HaSysTimeParseErr:
		return "时间解析错误"
	case HaSysParamErr:
		return "参数错误"
	//用户相关
	case HaUserNotLogin:
		return "用户未登录或登录已失效，请重新登录"
	}

	return "未知错误"
}
