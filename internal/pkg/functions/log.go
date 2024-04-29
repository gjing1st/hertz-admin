// Path: internal/pkg
// FileName: log.go
// Created by bestTeam
// Author: GJing
// Date: 2022/11/7$ 16:38$

package functions

import (
	backend "github.com/gjing1st/hertz-admin"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strconv"
)

// AddErrLog
// @description: 记录错误日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/7 16:40
// @success:
func AddErrLog(errMsg log.Fields) {
	//记录上一步调用者文件行号
	//go func() {
	_, file, lineNo, _ := runtime.Caller(1)
	errMsg["log_file"] = file + ":" + strconv.Itoa(lineNo)
	log.WithFields(errMsg).Error(backend.ProjectName)
	//}()
}

func AddWarnLog(errMsg log.Fields) {
	//记录上一步调用者文件行号
	//go func() {
	_, file, lineNo, _ := runtime.Caller(1)
	errMsg["log_file"] = file + ":" + strconv.Itoa(lineNo)
	log.WithFields(errMsg).Warn(backend.ProjectName)
	//}()
}

// AddInfoLog
// @description: 记录日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/25 15:59
// @success:
func AddInfoLog(fields log.Fields) {
	//go func() {
	_, file, lineNo, _ := runtime.Caller(1)
	fields["log_file"] = file + ":" + strconv.Itoa(lineNo)
	log.WithFields(fields).Info(backend.ProjectName)
	//}()
}

// AddDebugLog
// @description: debug记录
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/28 15:14
// @success:
func AddDebugLog(fields log.Fields) {
	//go func() {
	_, file, lineNo, _ := runtime.Caller(1)
	fields["log_file"] = file + ":" + strconv.Itoa(lineNo)
	log.WithFields(fields).Debug(backend.ProjectName)
	//}()
}
