// Path: internal/apiserver/controller
// FileName: sys.go
// Created by bestTeam
// Author: GJing
// Date: 2024/4/29$ 11:29$

package controller

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/response"
	"github.com/gjing1st/hertz-admin/internal/apiserver/service"
)

type SysController struct {
}

var sysService service.SysService

// ServerStatus
// @description: 设备运行状态
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/9 21:30
// @success:
func (slc *SysController) ServerStatus(ctx context.Context, c *app.RequestContext) {
	res, _ := sysService.ServerStatus()
	response.OkWithData(c, res)

}

// SysRunDate
// @description: 设备运行时间
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 11:19
// @success:
func (slc *SysController) SysRunDate(ctx context.Context, c *app.RequestContext) {
	res, errCode := configService.GetRunDate()
	if errCode != nil {
		response.Failed(c, errCode)
		return
	}
	response.OkWithData(c, res)

}
