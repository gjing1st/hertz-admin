// Path: internal/apiserver/model/dict
// FileName: config.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 9:47$

package dict

import (
	"time"
)

const (
	ConfigInitKey              = "initialized"                   //系统是否已初始化
	ConfigSysFirstStartDate    = "first_start_date"              //系统首次运行的时间
	ConfigSysBreakDate         = "sys_break_date"                //系统故障时间
	ConfigLoginType            = "login_type"                    //登录方式
	ConfigInitStep             = "init_step"                     //初始化步骤
	ConfigGuideStep            = "guide_step"                    //向导步骤
	ConfigVersion              = "version"                       //当前版本信息
	ConfigLatestVersion        = "latest_version"                //最新版本信息
	ConfigBackupTime           = "backup_time"                   //备份时间
	ConfigRestoreTime          = "restore_time"                  //恢复时间
	ConfigAutoUpdate           = "auto_update"                   //是否自动更新
	ConfigUpdateRange          = "update_range"                  //自动更新周期
	ConfigUpdateTime           = "update_time"                   //自动更新时间
	ConfigBackupPeriod         = "backup_period"                 //备份周期
	ConfigWhitelistStatus      = "whitelist_status"              //白名单管理
	ConfigPwdValidDate         = "pwd_valid_date"                //密码有效期
	PlatformConfigRate         = "platform_config_rate"          //数据抓取周期
	PlatformConfigProvince     = "platform_config_province"      //省
	PlatformConfigCity         = "platform_config_city"          //市
	PlatformConfigCSMPIP       = "platform_config_csmp_ip"       //数据抓取周期
	PlatformConfigAppid        = "platform_config_appid"         //数据抓取周期
	PlatformConfigAppSecret    = "platform_config_app_secret"    //数据抓取周期
	PlatformConfigTenantSecret = "platform_config_tenant_secret" //数据抓取周期
	PlatformConfigTenantKey    = "platform_config_tenant_key"    //数据抓取周期
)

const (
	LoginTypePasswd            = 1     //用户名口令
	LoginTypeFrontUKey         = 2     //前端UKey登录
	LoginTypeBackendUKey       = 3     //后端UKey登录
	VersionModeNormal          = 1     //发货版本
	VersionModeNormalLoginType = "1,2" //发货版本
	VersionModeCheck           = 2     //送检版本
)

// 是否已初始化完成
const (
	InitStepValueNot  = 0 //未配置
	InitStepValueDown = 1 //已完成配置
)

// 初始化步骤
const (
	InitStepUser    = 1 //步骤添加管理
	InitStepNetwork = 2 //步骤配置网络
	InitStepReset   = 3 //初始化重置
)

const (
	PwdValidDateForever    = iota + 1 //密码时效性-永久
	PwdValidDate30Day                 //30天
	PwdValidDateThreeMonth            //3个月
	PwdValidDateSixMonth              //6个月
)

func CheckPwdValidDate(policy int, pwdChangeDate time.Time) (pass bool) {
	switch policy {
	case PwdValidDate30Day:
		if time.Now().After(pwdChangeDate.AddDate(0, 0, 30)) {
			pass = true
		}
	case PwdValidDateThreeMonth:
		if time.Now().After(pwdChangeDate.AddDate(0, 3, 0)) {
			pass = true
		}
	case PwdValidDateSixMonth:
		if time.Now().After(pwdChangeDate.AddDate(0, 6, 0)) {
			pass = true
		}
	default:
		pass = false
	}

	return pass
}
