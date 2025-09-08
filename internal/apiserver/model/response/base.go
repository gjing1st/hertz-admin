// Path: internal/apiserver/model/response
// FileName: base.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/29$ 14:02$

package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/dict"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/database"
	"github.com/gjing1st/hertz-admin/pkg/errcode"
	"github.com/gjing1st/hertz-admin/pkg/utils"
)

// Response 响应格式
type Response struct {
	Code error       `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// PageResult 分页格式
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

func Result(code error, data interface{}, c *app.RequestContext) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		code.Error(),
	})
}

// OkWithData
// @params
// @Description 请求处理成功，返回响应200
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/3/31 19:02
func OkWithData(c *app.RequestContext, data interface{}) {
	Result(errcode.SuccessCode, data, c)
}

// Ok
// @Description 请求处理成功，直接返回200
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/3/31 20:00
func Ok(c *app.RequestContext) {
	Result(errcode.SuccessCode, nil, c)
}

// Failed
// @Description 请求失败，影响返回500
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/3/31 19:56
func Failed(c *app.RequestContext, code error) {
	res := Response{
		code,
		nil,
		code.Error(),
	}
	c.JSON(http.StatusInternalServerError, res)
}
func FailedWithData(c *app.RequestContext, code error, data interface{}, msg string) {
	res := Response{
		code,
		data,
		msg,
	}
	c.JSON(http.StatusInternalServerError, res)
}

// FailWithLog
// @description: 返回操作失败，并记录失败日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 16:20
// @success:
func FailWithLog(err error, content string, req interface{}, c *app.RequestContext) {
	code, ok := err.(error)
	if !ok {
		code = errcode.ErrCode
	}
	c.JSON(http.StatusInternalServerError, Response{
		code,
		nil,
		code.Error(),
	})
	RecordLog(content, req, c)
}

// FailWithDataLog
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/9 17:48
// @success:
func FailWithDataLog(data interface{}, err error, content string, req interface{}, c *app.RequestContext) {
	code, ok := err.(error)
	if !ok {
		code = errcode.ErrCode
	}
	msg := code.Error()
	if errors.Is(err, errcode.PwdErr) {
		res := data.(UserLogin)
		if res.RetryCount > 0 {
			msg = msg + strconv.Itoa(res.RetryCount)
		} else {
			msg = "密码输入错误次数过多，已锁定"
		}
	}
	c.JSON(http.StatusInternalServerError, Response{
		code,
		data,
		msg,
	})
	RecordLog(content, req, c)
}

// OkWithLog
// @description: 返回操作成功并记录操作日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 16:16
// @success:
func OkWithLog(content string, req interface{}, c *app.RequestContext) {
	Result(errcode.SuccessCode, nil, c)
	RecordLog(content, req, c)
}

// OkWithDataLog
// @description: 返回操作成功并记录操作日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 17:30
// @success:
func OkWithDataLog(data interface{}, content string, req interface{}, c *app.RequestContext) {
	Result(errcode.SuccessCode, data, c)
	RecordLog(content, req, c)

}

// ParamErr
// @Description 请求参数错误,返回400
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/3/31 19:58
func ParamErr(c *app.RequestContext) {
	c.JSON(http.StatusBadRequest, Response{
		errcode.HaSysParamErr,
		nil,
		errcode.HaSysParamErr.Error(),
	})
}

// Unauthorized
// @description: 未登录的错误
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 18:02
// @success:
func Unauthorized(err error, c *app.RequestContext) {
	code, ok := err.(error)
	if !ok {
		code = errcode.ErrCode
	}
	c.JSON(http.StatusUnauthorized, Response{
		code,
		nil,
		code.Error(),
	})
}

// Forbidden
// @description: 没有权限
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 18:05
// @success:
func Forbidden(err error, c *app.RequestContext) {
	code, ok := err.(error)
	if !ok {
		code = errcode.ErrCode
	}
	c.JSON(http.StatusForbidden, Response{
		code,
		nil,
		code.Error(),
	})
}

func RecordLog(content string, req interface{}, c *app.RequestContext) {
	go func() {
		var sysLog entity.SysLog
		sysLog.Result = dict.AdminLogResultOk
		sysLog.ClientIP = c.ClientIP()
		sysLog.Content = content
		userName, _ := c.Get("username")
		username := utils.String(userName)
		//if username == global.SuperAdmin {
		//	return
		//}
		sysLog.Username = username
		reqJson, _ := json.Marshal(req)
		sysLog.RequestData = string(reqJson)
		//sysLog.Address = ip.GetAddress(c.ClientIP())
		_ = database.SysLogDB{}.Create(nil, &sysLog)
	}()

}
