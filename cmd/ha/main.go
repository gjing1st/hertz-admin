// Path: cmd/ha
// FileName: main.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/28$ 21:03$

package main

import (
	"github.com/gjing1st/hertz-admin/internal/apiserver"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/database"
	"github.com/gjing1st/hertz-admin/internal/pkg/config"
)

// @title HertzAdmin
// @version 1.0
// @description This is a demo using Hertz.
// @contact.name gjing1st@gmail.com
// @contact.url http://tna.cn
// @contact.email gjing1st@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9681
// @BasePath /ha/v1
//
//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
func main() {
	//加载配置文件
	config.Init()
	//加载数据库驱动并初始化数据
	store.DB = database.GetDB()
	if store.DB != nil {
		db, _ := store.DB.DB()
		// 程序结束前关闭数据库链接
		defer db.Close()
	}
	//开启http服务
	apiserver.HttpStart()

}
