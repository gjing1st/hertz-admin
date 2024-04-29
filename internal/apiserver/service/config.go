// Path: internal/apiserver/service
// FileName: config.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 10:57$

package service

import (
	"encoding/json"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/dict"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/response"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/cache"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/database"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/database/initdata"
	"github.com/gjing1st/hertz-admin/internal/pkg/config"
	"github.com/gjing1st/hertz-admin/internal/pkg/functions"
	"github.com/gjing1st/hertz-admin/pkg/errcode"
	"github.com/gjing1st/hertz-admin/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type ConfigService struct {
}

var (
	gCache   = cache.Cache{}
	configDB = database.ConfigDB{}
)

// GetValueStr
// @description: 根据k获取v
// @param: k string
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 11:08
// @success:
func (cs *ConfigService) GetValueStr(k string) (v string, errCode error) {
	v, errCode = gCache.GetValueStr(k)
	if v == "" {
		vi, errCode1 := configDB.GetValue(k)
		if errCode1 != nil {
			errCode = errCode1
		}
		v = utils.String(vi)
		//存入缓存
		_ = cs.SetCacheValue(k, v)
	}
	return
}

// GetInitStep
// @description: 获取初始化步骤
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 10:01
// @success:
func (cs *ConfigService) GetInitStep() (res response.InitStepValue, errCode error) {
	v, errCode1 := cs.GetValueStr(dict.ConfigInitStep)
	if errCode1 != nil {
		errCode = errCode1
		return
	}
	err := json.Unmarshal([]byte(v), &res)
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "初始化步骤值转json错误", "err": err, "data": v, "v": res})
		errCode = errcode.HaSysJsonUnMarshalErr
		return
	}
	return
}

func (cs *ConfigService) Get() {

}

// SetCacheValue
// @description: 设置缓存
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 15:27
// @success:
func (cs *ConfigService) SetCacheValue(k, v interface{}) (errCode error) {
	errCode = gCache.RemoveSet(k, v)
	return
}

// SetValue
// @description: 设置k-v值，先持久化再更新缓存
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/15 14:01
// @success:
func (cs *ConfigService) SetValue(k string, v interface{}) (errCode error) {
	errCode = configDB.SetValue(k, v)
	if errCode == nil {
		_ = cs.SetCacheValue(k, v)
	}
	return
}

// GetRunDate
// @description: 获取系统运行时长
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 11:36
// @success:
func (cs *ConfigService) GetRunDate() (res response.SysRunDate, errCode error) {
	v, errCode1 := cs.GetValueStr(dict.ConfigSysBreakDate)
	if errCode1 != nil {
		errCode = errCode1
		return
	}
	//转换为time类型
	breakTime, err := time.ParseInLocation(time.DateTime, v, time.Local)
	if err != nil {
		errCode = errcode.SysTimeParseErr
		functions.AddErrLog(log.Fields{"msg": "breakTime时间转换出错", "err": err, "time": v})
		return
	}
	//当前时间与故障时间对比
	d := time.Now().Sub(breakTime)
	//day
	h := d / time.Hour //相差的小时数
	res.Minute = int(d % time.Hour / time.Minute)
	res.Hour = int(h % 24)
	res.Day = int(h / 24)
	return
}

// SysReset
// @description: 恢复出厂设置
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/30 13:48
// @success:
func (cs *ConfigService) SysReset() (err error) {
	//初始化状态
	//errCode = cs.SetInitStep(dict.InitStepReset)
	//if err != nil {
	//	return err
	//}
	err = store.DB.Transaction(func(tx *gorm.DB) error {
		//清理缓存
		store.GC.Purge()
		//删除管理员
		//err = userMysql.ResetUser(tx)
		//if err != nil {
		//	return err
		//}
		////清空日志记录
		//err = sysLogMysql.TruncateTable(tx)
		//if err != nil {
		//	return err
		//}
		////配置信息
		//err = configDB.TruncateTable(tx)
		//if err != nil {
		//	return err
		//}
		////清理白名单
		//err = whitelistDB.TruncateTable(tx)
		//if err != nil {
		//	return err
		//}
		//_ = util.Forbidden()
		var initConfig initdata.InitConfig
		err = initConfig.InitializeData(tx)
		if err != nil {
			return err
		}

		//TODO 删除其他数据
		return err
	})

	return
}

// VersionInfo
// @description: 获取当前版本信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/1/4 9:35
// @success:
func (cs *ConfigService) VersionInfo() (res response.VersionInfo) {
	res.Version, _ = cs.GetValueStr(dict.ConfigVersion)
	res.Manufacturer = config.Config.VersionInfo.Manufacturer
	res.Serial = config.Config.VersionInfo.Serial
	res.DeviceModel = config.Config.VersionInfo.Serial
	return
}
