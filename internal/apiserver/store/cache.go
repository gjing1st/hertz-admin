// Path: internal/apiserver/store/database
// FileName: cache.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 11:10$

package store

import (
	"github.com/bluele/gcache"
	"time"
)

var gc gcache.Cache

func Init() gcache.Cache {
	gc = gcache.New(200).
		ARC().
		Expiration(time.Hour * 8).
		Build()
	return gc
}

func GetCache() gcache.Cache {
	return Init()
}
