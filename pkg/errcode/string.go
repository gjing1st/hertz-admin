// Path: pkg/errcode
// FileName: string.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/28$ 21:44$

package errcode

func (e Err) Error() string {
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
		//用户相关
	case NoToken, TokenExpired:
		return "请重新登录"
	// 高级配置模块
	case UpdateFileErr, UpdateFileReadErr:
		return "升级包有误"
	case UpdateFileLoadErr:
		return "升级包镜像导入错误"
	case UpdateAssistErr:
		return "请求助手升级失败"
	// 系统运行中的缓存错误
	case SysCacheErr:
		return "缓存错误"
	case SysCacheGetErr:
		return "缓存获取错误"
	case SysCacheSetErr:
		return "缓存设置错误"
	// 系统运行中数据库出现的错误
	case SysCmdErr:
		return "执行宿主机cmd指令出错"
	case SysNetworkErr:
		return "网卡配置出错"
	// 系统运行中的错误
	case SysJsonMarshalErr:
		return "转json失败"
	case SysJsonUnMarshalErr:
		return "json解析失败"
	case SysTimeParseErr:
		return "时间转换错误"
	case SysSaveFileErr:
		return "保存文件错误"
	case UserNotFound:
		return "用户不存在"
	case PasswordErr:
		return "用户名或密码错误"
	case PwdErr:
		return "密码输入错误，剩余次数："
	case PwdMaxErr:
		return "密码错误次数过多，账号已锁定"
	case UKeyUserNotFound:
		return "用户不存在"
	case UKeyOpenErr, UKeyNotExist:
		return "请把UKey插入服务器"
	case UKeySerialErr:
		return "UKey序列号错误"
	case UKeyPwdErr:
		return "PIN输入错误，剩余次数："
	case MAXAdmin:
		return "管理员数量已达最大限制"
	case UKeyCertErr, UKeyParseErr, UKeyValidateErr, UKeyUsernameErr:
		return "新增失败"
	case UKeyAddUsernameErr:
		return "UKey用户名错误"
	case UserNameExist:
		return "当前用户已注册，新增失败"
	case DBErr, DBCreateErr, DBFindErr, DBUpdateErr, DBDeleteErr:
		return "db-系统错误"
	case OperationFail:
		return "操作失败"
	case ValueErr:
		return "数值错误"
	case IdCardNumberErr:
		return "身份证号错误"
	}

	return "未知错误"
}
