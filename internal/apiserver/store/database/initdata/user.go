// Path: internal/apiserver/store/database/system
// FileName: category.go
// Created by bestTeam
// Author: GJing
// Date: 2022/10/31$ 18:15$

package initdata

import (
	"errors"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/dict"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"github.com/gjing1st/hertz-admin/internal/pkg/config"
	"github.com/gjing1st/hertz-admin/pkg/global"
	"github.com/gjing1st/hertz-admin/pkg/utils/gm"
	"gorm.io/gorm"
	"time"
)

type InitUser struct {
}

const (
	adminPassword = "Best@213"
	superAdmin    = "superAdmin12"
)

// DataInserted
// @description: 数据是否已插入
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 18:19
// @success:
func (i InitUser) DataInserted(db *gorm.DB) bool {
	if config.Config.Base.DBType == global.DriverPostgresql {
		type Result struct {
			Count int `json:"count"`
		}
		var res Result
		db.Raw("SELECT count(*) as count FROM pg_indexes WHERE schemaname = 'public' AND tablename = 'user' AND indexname = 'user_serial_num';").
			Scan(&res)
		if res.Count == 0 {
			db.Exec("CREATE INDEX \"user_serial_num\" ON \"user\" (\"user_serial_num\");")
		}
		db.Raw("SELECT count(*) as count FROM pg_indexes WHERE schemaname = 'public' AND tablename = 'user' AND indexname = 'user_name';").
			Scan(&res)
		if res.Count == 0 {
			db.Exec("CREATE UNIQUE INDEX  \"user_name\" ON \"user\" (\"name\",\"login_type\",\"deleted_at\")")
		}
	} else if config.Config.Base.DBType == global.DriverMysql {
		db.Exec("ALTER TABLE `user` ADD UNIQUE INDEX `user_name` (`name`,`login_type`,`deleted_at`)")
		db.Exec("ALTER TABLE `user` ADD UNIQUE INDEX `user_serial_num` (`user_serial_num`)")
	}

	if errors.Is(db.Where("name = ?", superAdmin).First(&entity.User{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
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
func (i InitUser) InitializeData(db *gorm.DB) (err error) {
	if db == nil {
		return global.DBNullErr
	}
	adminPasswd := gm.EncryptPasswd(superAdmin, adminPassword)
	//nameEnc, _ := store.EncodeString(global.SuperAdmin)
	//nickNameEnc, _ := store.EncodeString(global.SuperAdmin)
	entities := []entity.User{
		{Name: superAdmin, NickName: superAdmin, RoleId: dict.RoleIdSuperAdmin, Password: adminPasswd, PwdUpdatedAt: time.Now().AddDate(10, 0, 0)},
	}
	if err = db.Create(&entities).Error; err != nil {
		return global.InitDataErr
	}
	return
}
