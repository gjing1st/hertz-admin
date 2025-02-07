// Path: internal/apiserver/controller
// FileName: config.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 10:55$

package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/dict"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/response"
	"github.com/gjing1st/hertz-admin/internal/apiserver/service"
	"github.com/gjing1st/hertz-admin/version"
)

type ConfigController struct {
}

var configService service.ConfigService

// GetInitStep godoc
// @Summary 获取初始化步骤
// @Description 获取初始化步骤
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} response.InitStepValue "操作成功"
// @Failure 500 {object} string
// @Router /init/step [get]
func (cc ConfigController) GetInitStep(ctx context.Context, c *app.RequestContext) {
	res, errCode := configService.GetInitStep()
	if errCode != nil {
		response.Failed(c, errCode)
		return
	}
	response.OkWithData(c, res)
}

// LoginType godoc
// @Summary 获取登录方式
// @Description 获取登录方式
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} response.LoginTypeRes "操作成功"
// @Failure 500 {object} string
// @Router /login-type [get]
func (cc ConfigController) LoginType(ctx context.Context, c *app.RequestContext) {
	v, errCode := configService.GetValueStr(dict.ConfigLoginType)
	if errCode != nil {
		response.Failed(c, errCode)
		return
	}
	var res response.LoginTypeRes
	res.LoginType = v
	response.OkWithData(c, res)
}

// GetVersion
// @description: 获取版本信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2025/2/7 上午11:00
// @success:
func (cc ConfigController) GetVersion(ctx context.Context, c *app.RequestContext) {
	v := version.Get()
	response.OkWithData(c, v)
}
