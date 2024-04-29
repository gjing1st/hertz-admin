// Path: internal/apiserver/store/mysql
// FileName: adminlog.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/28$ 14:56$

package database

import (
	"fmt"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/dict"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/request"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store"
	"github.com/gjing1st/hertz-admin/internal/pkg/functions"
	"github.com/gjing1st/hertz-admin/pkg/errcode"
	"github.com/gjing1st/hertz-admin/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type SysLogDB struct {
}

// Create
// @description: 创建管理员日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 14:59
// @success:
func (sl SysLogDB) Create(tx *gorm.DB, data *entity.SysLog) (errCode error) {
	if tx == nil {
		tx = store.DB
	}
	data.CheckData = sl.computeLogCheckData(data)
	err := tx.Create(&data).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "新增日志失败"})
		return errcode.DBCreateErr
	}
	return

}

// List
// @description: 日志列表查询
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/9 11:30
// @success:
func (sl SysLogDB) List(req *request.SysLogList) (logs []entity.SysLog, total int64, errCode error) {
	//db := store.DB.Model(&entity.SysLog{}).Where("category = ?", req.Category)
	db := store.DB.Model(&entity.SysLog{})
	if req.Keyword != "" {
		if strings.Contains(req.Keyword, "成功") {
			i := strings.Index(req.Keyword, "成功")
			req.Keyword = req.Keyword[:i]
			db.Where("username like ? or content like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%").Where("result=?", dict.AdminLogResultOk)

		} else if strings.Contains(req.Keyword, "失败") {
			i := strings.Index(req.Keyword, "失败")
			req.Keyword = req.Keyword[:i]
			db.Where("username like ? or content like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%").Where("result=?", dict.AdminLogResultFail)

		} else {
			db.Where("username like ? or content like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")

		}
	}
	if !req.StartDate.IsZero() {
		db.Where("created_at > ?", req.StartDate)
	}

	if !req.EndDate.IsZero() {
		//db.Where("DATE_FORMAT(`created_at`,'%Y-%m-%d') <= ?", req.StartDate)
		db.Where("created_at  < ?", req.EndDate)
	}
	//总数使用缓存，避免并发CPU占用
	cacheCount, err := store.GC.Get("sysLogCount")
	if err == nil {
		total = utils.Int64(cacheCount)
	} else {
		_ = db.Count(&total).Error
		_ = store.GC.SetWithExpire("sysLogCount", total, time.Second)
	}
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&logs).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询策略列表失败"})
		errCode = errcode.UserNotFound
	}
	for _, v := range logs {
		hashStr := sl.computeLogCheckData(&v)
		if v.CheckData != hashStr {
			functions.AddWarnLog(log.Fields{"msg": "该条数据完整性校验不通过，存在篡改嫌疑", "logData": v})
		}
	}
	return
}

func (sl SysLogDB) Export(req *request.SysLogExport) (logs []entity.SysLog, total int64, errCode error) {
	db := store.DB.Model(&entity.SysLog{}).Where("category = ?", req.Category)
	if req.Keyword != "" {
		if strings.Contains(req.Keyword, "成功") {
			i := strings.Index(req.Keyword, "成功")
			req.Keyword = req.Keyword[:i]
			db.Where("username like ? or content like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%").Where("result=?", dict.AdminLogResultOk)

		} else if strings.Contains(req.Keyword, "失败") {
			i := strings.Index(req.Keyword, "失败")
			req.Keyword = req.Keyword[:i]
			db.Where("username like ? or content like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%").Where("result=?", dict.AdminLogResultFail)

		} else {
			db.Where("username like ? or content like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")

		}
	}
	if !req.StartDate.IsZero() {
		db.Where("created_at > ?", req.StartDate)
	}

	if !req.EndDate.IsZero() {
		//db.Where("DATE_FORMAT(`created_at`,'%Y-%m-%d') <= ?", req.StartDate)
		db.Where("created_at  < ?", req.EndDate)
	}
	//总数使用缓存，避免并发CPU占用
	cacheCount, err := store.GC.Get("sysLogCount")
	if err == nil {
		total = utils.Int64(cacheCount)
	} else {
		_ = db.Count(&total).Error
		_ = store.GC.SetWithExpire("sysLogCount", total, time.Second)
	}
	err = db.Order("id desc").Find(&logs).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询策略列表失败"})
		errCode = errcode.UserNotFound
	}
	//for _, v := range logs {
	//	b := sl.hmacVerify(&v)
	//	if !b {
	//		functions.AddWarnLog(log.Fields{"msg": "该条数据完整性校验不通过，存在篡改嫌疑", "logData": v})
	//	}
	//}
	return
}

// TruncateTable
// @description: 清空日志表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/24 10:51
// @success:
func (sl SysLogDB) TruncateTable(tx *gorm.DB) (errCode error) {
	if tx == nil {
		tx = store.DB
	}
	err := tx.Exec("TRUNCATE TABLE " + entity.SysLog{}.TableName()).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql重置日志失败"})
		return errcode.DBDeleteErr
	}
	return
}

func (sl SysLogDB) computeLogCheckData(l *entity.SysLog) string {
	data := l.Username + l.Content + strconv.Itoa(l.Result)
	//hashStr, _ := store.StorageHmac(data)
	hashStr := fmt.Sprintf("%x", store.ComputeCheckData([]byte(data)))
	return hashStr
}

//func (sl SysLogDB) hmacVerify(l *entity.SysLog) bool {
//	value := l.Username + l.Content + strconv.Itoa(l.Result)
//	res, _ := store.StorageHMACVerify(l.CheckData, value)
//	return res
//
//}
