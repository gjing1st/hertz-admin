// Path: internal/apiserver/store/database/system
// FileName: category.go
// Created by bestTeam
// Author: GJing
// Date: 2022/10/31$ 18:15$

package initdata

import (
	"encoding/json"
	"errors"
	backend "github.com/gjing1st/hertz-admin"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/dict"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"github.com/gjing1st/hertz-admin/internal/pkg/config"
	"github.com/gjing1st/hertz-admin/pkg/global"
	"github.com/gjing1st/hertz-admin/pkg/utils"
	"gorm.io/gorm"
	"io/ioutil"
	"time"
)

type InitConfig struct {
}

// DataInserted
// @description: 数据是否已插入
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 18:19
// @success:
func (i InitConfig) DataInserted(db *gorm.DB) bool {
	if config.Config.Base.DBType == global.DriverPostgresql {
		type Result struct {
			Count int `json:"count"`
		}
		var res Result
		db.Raw("SELECT count(*) as count FROM pg_indexes WHERE schemaname = 'public' AND tablename = 'config' AND indexname = 'key_name';").
			Scan(&res)
		if res.Count == 0 {
			db.Exec("CREATE UNIQUE INDEX \"key_name\" ON \"config\" (\"name\")")
		}
	} else if config.Config.Base.DBType == global.DriverMysql {
		db.Exec("ALTER TABLE `config` ADD UNIQUE INDEX `key_name`(`name`)")
	}
	if errors.Is(db.Where("name = ?", dict.ConfigSysFirstStartDate).First(&entity.Config{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

// InitializeData
// @description: 初始化数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 17:26
// @success:
func (i InitConfig) InitializeData(db *gorm.DB) (err error) {
	if db == nil {
		return global.DBNullErr
	}
	//initStep := `{"user":0,"network":0}`
	initStep := `{"user":1,"network":1}`
	//读取版本信息
	var version = struct {
		Version string `json:"version"`
	}{}

	d, err := ioutil.ReadFile("./config/version.json")
	if err == nil {
		err = json.Unmarshal(d, &version)
	} else {
		version.Version = backend.Version
	}
	loginType := dict.LoginTypePasswd
	//发货版本送检版本
	//if config.Config.VersionModel == dict.VersionModeCheck {
	//	loginType = strconv.Itoa(dict.LoginTypeBackendUKey)
	//}
	//向导步骤
	//var step2 dict.GuideStepValue
	//guideStep, _ := json.Marshal(step2)
	entities := []entity.Config{
		{Name: dict.ConfigInitKey, Value: "false"},
		{Name: dict.ConfigSysFirstStartDate, Value: time.Now().Format(time.DateTime)},
		{Name: dict.ConfigSysBreakDate, Value: time.Now().Format(time.DateTime)},
		{Name: dict.ConfigLoginType, Value: loginType},
		{Name: dict.ConfigInitStep, Value: initStep},
		//{Name: dict.ConfigGuideStep, Value: string(guideStep)},
		{Name: dict.ConfigVersion, Value: version.Version},
		{Name: dict.ConfigLatestVersion},
		{Name: dict.ConfigBackupTime},
		{Name: dict.ConfigRestoreTime},
		{Name: dict.ConfigAutoUpdate},
		{Name: dict.ConfigUpdateRange},
		{Name: dict.ConfigUpdateTime},
		{Name: dict.ConfigBackupPeriod, Value: utils.String(0)},
		{Name: dict.ConfigWhitelistStatus, Value: utils.String(dict.StatusForbidden)},
		{Name: dict.ConfigPwdValidDate, Value: utils.String(0)},
		{Name: dict.PlatformConfigRate, Value: utils.String(10)},
		{Name: dict.PlatformConfigProvince, Value: "山东省"},
		{Name: dict.PlatformConfigCity, Value: "济南市"},
	}
	if err = db.Create(&entities).Error; err != nil {
		return global.InitDataErr
	}
	return
}

// Update
// @description: 更新数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/10 16:50
// @success:
func (i InitConfig) Update(db *gorm.DB) error {
	return db.Model(&entity.Config{}).Where("name = ?", dict.ConfigSysBreakDate).Update("value", time.Now().Format(time.DateTime)).Error
}
