package _map

import (
	"sync"
	"time"
)

var CacheMap sync.Map

// Set
//	@description:	内存变量过期 类redis
//	@param:			key 变量名
//	@param:			value 变量值
//	@param:			exp 过期时间
//	@author:		GJing
//	@email:			guojing@tna.cn
//	@date:			2022/7/27 17:20
//	@success:
func Set(key, value interface{}, exp time.Duration) {
	CacheMap.Store(key, value)
	time.AfterFunc(exp, func() {
		CacheMap.Delete(key)
	})
}