// Path: internal/apiserver/store/cache
// FileName: gcache.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/4/7$ 23:35$

package cache

import (
	"github.com/bluele/gcache"
	"time"
)

var gc gcache.Cache

func init() {
	gc = gcache.New(200).
		ARC().
		Expiration(time.Hour * 8).
		Build()
}

func GetCache() gcache.Cache {
	return gc
}
