// Path: internal/apiserver/store/mysql
// FileName: user.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 16:00$

package database

import (
	"errors"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/request"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store"
	"github.com/gjing1st/hertz-admin/internal/pkg/functions"
	errcode "github.com/gjing1st/hertz-admin/pkg/errcode"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type UserDB struct {
}

// Create
// @description: mysql存储
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 16:00
// @success:
func (um UserDB) Create(tx *gorm.DB, user *entity.User) (id uint, errCode error) {
	if tx == nil {
		tx = store.DB
	}
	if user.PwdUpdatedAt.IsZero() {
		user.PwdUpdatedAt = time.Now()
	}
	//user.NickName, _ = store.EncodeString(user.NickName)
	//user.Name, _ = store.EncodeString(user.Name)
	//user.CheckData = um.computeUserCheckData(user)
	err := tx.Create(&user).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql创建管理员失败"})
		return id, errcode.DBCreateErr
	}
	return user.ID, nil
}

// GetByName
// @description: 通过用户名查询用户信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 17:10
// @success:
func (um UserDB) GetByName(name string) (user *entity.User, errCode error) {
	//nameEnc, _ := store.EncodeString(name)
	err := store.DB.Where("name = ?", name).First(&user).Error
	if err != nil {
		functions.AddInfoLog(log.Fields{"err": err, "msg": "mysql创建管理员失败"})
		if errors.Is(err, errcode.ErrRecordNotFound) {
			return user, errcode.UserNotFound
		}
		return user, errcode.DBFindErr
	}
	//user.NickName, _ = store.DecodeString(user.NickName)
	//user.Name, _ = store.DecodeString(user.Name)
	//hashStr := um.computeUserCheckData(user)
	//if hashStr != user.CheckData {
	//	functions.AddWarnLog(log.Fields{"msg": "该条数据完整性校验不通过，存在篡改嫌疑", "user": user})
	//}
	return
}

// GetByNameLoginType
// @Description 根据用户名和登录方式查找用户信息
// @params name string 用户名
// @params loginType int 用户名
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/6/6 14:54
func (um UserDB) GetByNameLoginType(name string, loginType int) (user *entity.User, errCode error) {
	//err := store.DB.Where("name = ? and login_type = ?", name, loginType).First(&user).Error
	err := store.DB.Where("name = ? ", name).First(&user).Error
	if err != nil {
		functions.AddWarnLog(log.Fields{"err": err, "msg": "mysql查询管理员失败"})
		if errors.Is(err, errcode.ErrRecordNotFound) {
			return user, nil
		}
		return user, errcode.DBFindErr
	}
	//user.NickName, _ = store.DecodeString(user.NickName)
	//user.Name, _ = store.DecodeString(user.Name)
	//hashStr := um.computeUserCheckData(user)
	//if hashStr != user.CheckData {
	//	functions.AddWarnLog(log.Fields{"msg": "该条数据完整性校验不通过，存在篡改嫌疑", "user": user})
	//}
	return
}

// AddErrNum
// @Description 密码错误次数+1
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/11/6 16:28
func (um UserDB) AddErrNum(id int) (err error) {
	err = store.DB.Model(&entity.User{}).Where("id = ?", id).UpdateColumn("err_num", gorm.Expr("err_num + ?", 1)).Error
	if err != nil {
		functions.AddWarnLog(log.Fields{"err": err, "msg": "mysql查询管理员密码错误次数+1失败"})
		return errcode.DBUpdateErr
	}
	return
}

func (um UserDB) ClearErrNum(id int) (err error) {
	err = store.DB.Model(&entity.User{}).Where("id = ?", id).UpdateColumn("err_num", 0).Error
	if err != nil {
		functions.AddWarnLog(log.Fields{"err": err, "msg": "mysql查询管理员密码错误次数清除失败"})
		return errcode.DBUpdateErr
	}
	return
}

// UpdateToken
// @description: 修改数据库用户token
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 18:14
// @success:
func (um UserDB) UpdateToken(id uint, token string) (errCode error) {
	err := store.DB.Model(&entity.User{}).Where("id = ?", id).Update("token", token).Error
	if err != nil {
		errCode = errcode.DBUpdateErr
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql修改token失败", "id": id, "token": token})
	}
	return
}

// List
// @description: 获取用户列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:31
// @success:
func (um UserDB) List(req *request.UserList) (users []entity.User, total int64, errCode error) {
	db := store.DB.Model(&entity.User{})
	if req.Keyword != "" {
		db.Where("name like ?", "%"+req.Keyword+"%")
	}
	if len(req.RoleId) > 0 {
		db.Where("role_id in ?", req.RoleId)
	}
	err := db.Count(&total).Error
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&users).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询策略列表失败"})
		errCode = errcode.UserNotFound
	}
	//for _, user := range users {
	//users[i].NickName, _ = store.DecodeString(user.NickName)
	//users[i].Name, _ = store.DecodeString(user.Name)
	//hashStr := um.computeUserCheckData(&user)
	//if user.CheckData != hashStr {
	//	functions.AddWarnLog(log.Fields{"msg": "该条数据完整性校验不通过，存在篡改嫌疑", "user": user})
	//}
	//}
	return
}

// GetByNameAndSerialNum
// @description: 查询ukey管理员
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 15:36
// @success:
func (um UserDB) GetByNameAndSerialNum(name, serialNum string) (user *entity.User, errCode error) {
	//nameEnc, _ := store.EncodeString(name)
	err := store.DB.Where("name = ? and user_serial_num = ?", name, serialNum).First(&user).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询ukey管理员失败"})
		if err == errcode.ErrRecordNotFound {
			return user, errcode.UserNotFound
		}
		return user, errcode.DBFindErr
	}
	//user.NickName, _ = store.DecodeString(user.NickName)
	//user.Name, _ = store.DecodeString(user.Name)
	//hashStr := um.computeUserCheckData(user)
	//if hashStr != user.CheckData {
	//	functions.AddWarnLog(log.Fields{"msg": "该条数据完整性校验不通过，存在篡改嫌疑", "user": user})
	//}
	return
}

// ResetUser
// @description: 删除管理员意外的其他用户
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 18:38
// @success:
func (um UserDB) ResetUser(tx *gorm.DB) (errCode error) {
	if tx == nil {
		tx = store.DB
	}
	//永久删除
	err := tx.Unscoped().Delete(&entity.User{}).Error
	//err := tx.Exec("TRUNCATE TABLE " + entity.User{}.TableName()).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql重置管理员失败"})
		return errcode.DBDeleteErr
	}
	return
}

// DeleteById
// @description: 通过id删除用户
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/30 11:28
// @success:
func (um UserDB) DeleteById(tx *gorm.DB, userid int) (errCode error) {
	if tx == nil {
		tx = store.DB
	}
	err := tx.Where("id = ?", userid).Delete(&entity.User{}).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql重置管理员失败"})
		return errcode.DBDeleteErr
	}
	return
}

// Count
// @Description 所有用户综合
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/6/7 10:44
func (um UserDB) Count() (total int64, err error) {
	err = store.DB.Model(&entity.User{}).Count(&total).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql重置管理员失败"})
		err = errcode.DBFindErr
		return
	}
	return
}

// Save
// @Description
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/6/8 9:49
func (um UserDB) Save(user *entity.User) (err error) {
	//user.CheckData = um.computeUserCheckData(user)
	if user.PwdUpdatedAt.IsZero() {
		user.PwdUpdatedAt = time.Now()
	}
	//user.NickName, _ = store.EncodeString(user.NickName)
	//user.Name, _ = store.EncodeString(user.Name)
	err = store.DB.Save(&user).Error
	if err != nil {
		if errcode.IsDuplicateKeyError(err) {
			err = errcode.UserNameExist
			return
		}
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql保存管理员失败"})
		err = errcode.DBFindErr
		return
	}
	return
}

func (um UserDB) UpdatePwd(user *entity.User) (err error) {
	//user.CheckData = um.computeUserCheckData(user)
	if user.PwdUpdatedAt.IsZero() {
		user.PwdUpdatedAt = time.Now()
	}
	//user.NickName, _ = store.EncodeString(user.NickName)
	//user.Name, _ = store.EncodeString(user.Name)
	err = store.DB.Save(&user).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql保存管理员失败"})
		err = errcode.DBFindErr
		return
	}
	return
}

//func (um UserDB) computeUserCheckData(user *entity.User) string {
//	if !config.Config.Base.EnableIntegrity {
//		return ""
//	}
//	data := fmt.Sprintf("%x%s%x", user.RoleId, user.UserSerialNum, user.UserPub)
//	hashStr := fmt.Sprintf("%x", store.ComputeCheckData([]byte(data)))
//	return hashStr
//}

func (um UserDB) GetByPhone(phone string) (user *entity.User, errCode error) {
	err := store.DB.Where("phone = ? ", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, errcode.ErrRecordNotFound) {
			return user, nil
		}
		functions.AddWarnLog(log.Fields{"err": err, "msg": "查询" + phone + "用户失败"})
		return user, errcode.DBFindErr
	}
	return
}
