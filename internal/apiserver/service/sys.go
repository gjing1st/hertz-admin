// Path: internal/apiserver/service
// FileName: sys.go
// Created by bestTeam
// Author: GJing
// Date: 2024/4/29$ 11:31$

package service

import (
	"sync"

	"github.com/gjing1st/hertz-admin/internal/apiserver/model/dict"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/response"
	"github.com/gjing1st/hertz-admin/internal/pkg/functions"
	"github.com/gjing1st/hertz-admin/pkg/errcode"
	"github.com/gjing1st/hertz-admin/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type SysService struct {
}

// ServerStatus
// @description: 设备运行状态
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/9 20:29
// @success:
func (ss *SysService) ServerStatus() (res response.ServerStatus, errCode error) {
	res.ServiceStatus = dict.ServiceStatusInit
	res.RunStatus = dict.RunStatusAbnormal
	var wg sync.WaitGroup
	//TODO 需要在此补充需要检查的设备运行状态
	serviceStatus := make(chan bool, 1)
	runStatus := make(chan bool, 3)
	var err error
	wg.Go(func() {
		if err != nil {
			serviceStatus <- false
			return
		}
		serviceStatus <- true
	})
	wg.Go(func() {
		//TODO
		if err != nil {
			runStatus <- false
			return
		}
		runStatus <- true
	})
	wg.Go(func() {
		//TODO
		if err != nil {
			runStatus <- false
			return
		}
		runStatus <- true
	})
	wg.Go(func() {
		//TODO
		if err != nil {
			runStatus <- false
			return
		}
		runStatus <- true
	})

	wg.Wait()
	//运行状态
	//runStatusRes := <-runStatus
	//if runStatusRes {
	//	res.RunStatus = dict.RunStatusNormal
	//}
	//服务状态
	serviceStatusRes := <-serviceStatus
	if serviceStatusRes {
		res.ServiceStatus = dict.ServiceStatusReady
	}
	//运行状态
	for i := 0; i < 3; i++ {
		status := <-runStatus

		if status == false {
			//有未完成的
			return
		}
	}
	res.RunStatus = dict.RunStatusNormal
	return
}

// Reboot
// @description: 服务器重启
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/11 16:30
// @success:
func (ss *SysService) Reboot() (errCode error) {
	err := utils.DockerRunCommand("reboot")
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "服务器重启指令执行失败", "err": err})
		errCode = errcode.SysCmdErr
	}
	return
}
