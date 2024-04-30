// Path: internal/apiserver/router/v1
// FileName: sys.go
// Created by bestTeam
// Author: GJing
// Date: 2024/4/29$ 10:58$

package v1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/gjing1st/hertz-admin/internal/apiserver/controller"
)

func initSys(r *route.RouterGroup) {
	api := r.Group("sys")
	sysController := controller.SysController{}
	api.GET("run", sysController.SysRunDate)      //系统运行时长
	api.GET("status", sysController.ServerStatus) //系统运行状态

}
