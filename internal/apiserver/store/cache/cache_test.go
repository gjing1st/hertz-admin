// Path: internal/apiserver/store/cache
// FileName: cache_test.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 11:26$

package cache

import (
	"fmt"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store"
	"testing"
	"time"
)

var configCache Cache

func TestSet(t *testing.T) {
	configCache.RemoveSet("k", "12345")
	v, err := configCache.GetValueStr("k1")
	fmt.Println("v=", v)
	fmt.Println(err)
}

func TestTTl(t *testing.T) {
	store.GC.SetWithExpire("a", "a.value", time.Second*10)
	a, _ := store.GC.Get("a")
	fmt.Println("a=", a)
	go func() {
		for {
			time.Sleep(time.Second)
			a2, _ := store.GC.Get("a")
			fmt.Println("a2====", a2)
		}
	}()
	store.GC.Remove("a")
	time.Sleep(time.Second * 10)
	a1, _ := store.GC.Get("a")
	fmt.Println("a1====", a1)
	time.Sleep(time.Second * 10)
}
