// Path: internal/apiserver/store/cache
// FileName: config.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 11:09$

package cache

import (
	"errors"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store"
	"github.com/gjing1st/hertz-admin/internal/pkg/functions"
	"github.com/gjing1st/hertz-admin/pkg/errcode"
	"github.com/gjing1st/hertz-admin/pkg/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

type Cache struct {
}

// GetValueStr
// @description: 从缓存中获取value
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 11:20
// @success:
func (cc *Cache) GetValueStr(k interface{}) (v string, errCode error) {
	value, err := store.GC.Get(k)
	if err != nil {
		errCode = errcode.SysCacheGetErr
		functions.AddWarnLog(log.Fields{"err": err, "msg": "cache获取value错误", "key": k})
		if err == errcode.ErrKeyNotFound {
			errCode = nil
		}
	}
	v = utils.String(value)
	return
}

// RemoveSet
// @description: 删除旧缓存设置新缓存,该操作不适合高并发场景，会导致数据不一致。如果需要，自行加锁
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 11:24
// @success:
func (cc *Cache) RemoveSet(k, v interface{}) (errCode error) {
	store.GC.Remove(k)
	err := store.GC.SetWithExpire(k, v, time.Hour*8)
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "cache缓存key,value错误", "key": k, "value": v})
		errCode = errcode.SysCacheSetErr
	}
	return
}

// Get
// @description: 获取缓存中k对应的值
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 18:00
// @success:
func (cc *Cache) Get(k interface{}) (v interface{}, errCode error) {
	value, err := store.GC.Get(k)
	if err != nil {
		errCode = errcode.SysCacheGetErr
		functions.AddWarnLog(log.Fields{"err": err, "msg": "cache获取value错误", "key": k})
		if errors.Is(err, errcode.ErrKeyNotFound) {
			errCode = nil
		}
	}
	v = value
	return
}

// Remove
// @description: 删除缓存
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 18:10
// @success:
func (cc *Cache) Remove(k interface{}) bool {
	return store.GC.Remove(k)
}
