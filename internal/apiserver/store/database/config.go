// Path: internal/apiserver/store/mysql
// FileName: config.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 10:59$

package database

import (
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/dict"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store"
	"github.com/gjing1st/hertz-admin/internal/pkg/functions"
	errcode "github.com/gjing1st/hertz-admin/pkg/errcode"
	"github.com/gjing1st/hertz-admin/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ConfigDB struct {
}

// GetValue
// @description: 根据k获取v
// @param: k string
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 11:06
// @success:
func (cs ConfigDB) GetValue(k string) (v interface{}, err error) {
	var conf entity.Config
	err = store.DB.Where("name = ?", k).Where("status = ?", dict.StatusEnable).First(&conf).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			conf.Name = k
			err = store.DB.Model(&entity.Config{}).Save(&conf).Error
			if err != nil {
				functions.AddErrLog(log.Fields{"err": err, "msg": "mysql保存config数据错误", "key": k})
			}

		} else {
			functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询value错误", "key": k})
			err = errcode.DBFindErr
		}
		return "", err
	}
	return conf.Value, nil
}

// SetValue
// @description: 修改值
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 10:57
// @success:
func (cs ConfigDB) SetValue(k string, v interface{}) (err error) {
	conf := &entity.Config{
		Name:  k,
		Value: v,
	}
	vs := utils.String(v)
	err = store.DB.Model(&entity.Config{}).Where("name = ?", k).Update("value", vs).Error
	if err != nil {
		err = store.DB.Model(&entity.Config{}).Save(conf).Error
		if err != nil {
			functions.AddErrLog(log.Fields{"err": err, "msg": "mysql修改config表错误", "key": k})
			return errcode.DBUpdateErr
		}

	}

	return
}

// TruncateTable
// @Description 清空配置信息
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/6/9 9:54
func (cs ConfigDB) TruncateTable(tx *gorm.DB) (errCode error) {
	if tx == nil {
		tx = store.DB
	}
	err := tx.Exec("TRUNCATE TABLE " + entity.Config{}.TableName()).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql重置日志失败"})
		return errcode.DBDeleteErr
	}
	return
}
