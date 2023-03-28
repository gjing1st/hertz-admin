// Path: cmd/ha
// FileName: main.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/28$ 21:03$

package main

import (
	"github.com/gjing1st/hertz-admin/internal/apiserver"
	"github.com/gjing1st/hertz-admin/internal/pkg/config"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
func main() {
	//加载配置文件
	config.Init()
	//加载数据库驱动并初始化数据
	//store.DB = database.GetDB()
	//if store.DB != nil {
	//	db, _ := store.DB.DB()
	//	// 程序结束前关闭数据库链接
	//	defer db.Close()
	//}
	//开启http服务
	apiserver.HttpStart()

}
