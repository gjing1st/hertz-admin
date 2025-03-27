package controller

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/response"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/database"
	"github.com/gjing1st/hertz-admin/pkg/errcode"
	"github.com/gjing1st/hertz-admin/pkg/utils"
	"strconv"
)

type BaseController struct {
}

// RecordLog
// @description: 使用协程记录基础请求日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2024/11/15 下午2:35
// @success:
func RecordLog(ctx context.Context, c *app.RequestContext) {
	// c相关参数获取要写在go协程外层，不然hertz路由会匹配不到
	clientIp := c.ClientIP()
	go func() {
		content := ctx.Value("content")
		reqResult := ctx.Value("reqResult")
		var sysLog entity.SysLog
		sysLog.Result = utils.Int(reqResult)
		sysLog.ClientIP = clientIp
		sysLog.Content = utils.String(content)
		_ = database.SysLogDB{}.Create(nil, &sysLog)
		<-ctx.Done()

	}()
}

// DetailRecordLog
// @description: 记录详细请求日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2024/11/15 下午2:40
// @success:
func DetailRecordLog(ctx context.Context, c *app.RequestContext) {
	content := ctx.Value("content")
	req := ctx.Value("req")
	res := ctx.Value("res")
	reqResult := ctx.Value("reqResult")
	var sysLog entity.SysLog
	sysLog.Result = utils.Int(reqResult)
	sysLog.ClientIP = c.ClientIP()
	sysLog.Content = utils.String(content)
	userName, _ := c.Get("username")
	username := utils.String(userName)
	if len(username) == 0 {
		username = utils.String(ctx.Value("username"))
	}

	sysLog.Username = username
	sysLog.RequestData = utils.String(req)
	sysLog.ResultData = utils.String(res)
	_ = database.SysLogDB{}.Create(nil, &sysLog)

}

// GoDetailRecordLog
// @description: 使用go协程记录详细请求日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2024/11/15 下午2:42
// @success:
func GoDetailRecordLog(ctx context.Context, c *app.RequestContext) {
	clientIp := c.ClientIP()
	userName, _ := c.Get("username")
	go func() {
		content := ctx.Value("content")
		req := ctx.Value("req")
		res := ctx.Value("res")
		reqResult := ctx.Value("reqResult")
		var sysLog entity.SysLog
		sysLog.Result = utils.Int(reqResult)
		sysLog.ClientIP = clientIp
		sysLog.Content = utils.String(content)
		username := utils.String(userName)
		if len(username) == 0 {
			username = utils.String(ctx.Value("username"))
		}
		sysLog.Username = username
		sysLog.RequestData = utils.String(req)
		sysLog.ResultData = utils.String(res)
		_ = database.SysLogDB{}.Create(nil, &sysLog)
		<-ctx.Done()
	}()

}

const (
	LogTypeNormal = iota
	LogTypeDetail
	LogTypeGoDetail
	LogTypeWechat
)

// OkWithLog
// @description: 返回操作成功并记录操作日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 16:16
// @success:
func (bc BaseController) OkWithLog(ctx context.Context, c *app.RequestContext) {
	data := ctx.Value("res")
	if data != nil {
		response.Result(errcode.SuccessCode, data, c)
	} else {
		response.Result(errcode.SuccessCode, nil, c)
	}
	//按类型记录日志
	logType := ctx.Value("logType")
	switch logType {
	case LogTypeDetail:
		DetailRecordLog(ctx, c)
	case LogTypeGoDetail:
		GoDetailRecordLog(ctx, c)
	default:
		RecordLog(ctx, c)
	}
}
func (bc BaseController) OkWithData(c *app.RequestContext, res interface{}) {
	response.Result(errcode.SuccessCode, res, c)

}

// FailWithLog
// @description: 请求失败，携带数据返回
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2024/11/15 下午2:45
// @success:
func (bc BaseController) FailWithLog(ctx context.Context, c *app.RequestContext) {
	err := ctx.Value("err")
	code, ok := err.(error)
	if !ok {
		code = errcode.ErrCode
	}
	msg := code.Error()
	data := ctx.Value("res")

	if data != nil {
		if errors.Is(code, errcode.PwdErr) {
			res := data.(response.UserLogin)
			if res.RetryCount > 0 {
				msg = msg + strconv.Itoa(res.RetryCount)
			} else {
				msg = "密码输入错误次数过多，已锁定"
			}
		}
		response.FailedWithData(c, code, data, msg)
	} else {
		response.Failed(c, code)
	}

	//记录日志
	logType := ctx.Value("logType")
	switch logType {
	case LogTypeDetail:
		DetailRecordLog(ctx, c)
	case LogTypeGoDetail:
		GoDetailRecordLog(ctx, c)
	default:
		RecordLog(ctx, c)
	}
}

func (bc BaseController) Fail(code error, c *app.RequestContext) {
	response.Failed(c, code)
}

func (bc BaseController) Ok(c *app.RequestContext) {
	response.Result(errcode.SuccessCode, nil, c)
}

// ParamErr
// @Description 请求参数错误,返回400
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/3/31 19:58
func (bc BaseController) ParamErr(ctx context.Context, c *app.RequestContext) {

	response.ParamErr(c)
	//记录日志
	logType := ctx.Value("logType")
	switch logType {
	case LogTypeDetail:
		DetailRecordLog(ctx, c)
	case LogTypeGoDetail:
		GoDetailRecordLog(ctx, c)
	default:
		RecordLog(ctx, c)
	}
}
