// Path: internal/apiserver/router/v1
// FileName: without.go
// Created by bestTeam
// Author: GJing
// Date: 2024/4/28$ 10:12$

package v1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/gjing1st/hertz-admin/internal/apiserver/controller"
)

func initWithoutConfigRouter(apiV1 *route.RouterGroup) {
	configController := controller.ConfigController{}
	apiV1.GET("init/step", configController.GetInitStep) //初始化状态步骤
	apiV1.GET("login-type", configController.LoginType)  //登录方式
	apiV1.GET("version", configController.GetVersion)    //版本信息

}
