// Path: internal/apiserver/model/response
// FileName: base.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/29$ 14:02$

package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gjing1st/hertz-admin/pkg/errcode"
	"net/http"
)

// Response 响应格式
type Response struct {
	Code errcode.Err `json:"code"`
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

// OkWithData
// @params
// @Description 请求处理成功，返回响应200
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/3/31 19:02
func OkWithData(c *app.RequestContext, data interface{}) {
	//res := Response{
	//	errcode.SuccessCode,
	//	data,
	//	errcode.SuccessCode.String(),
	//}
	//resByte, err := sonic.Marshal(res)
	//if err != nil {
	//	Fail(c, errcode.HaSysJsonMarshalErr)
	//	return
	//}
	//c.Response.SetStatusCode(http.StatusOK)
	//c.Response.SetBody(resByte)
	c.JSON(http.StatusOK, Response{
		errcode.SuccessCode,
		data,
		errcode.SuccessCode.String(),
	})
}

// Ok
// @Description 请求处理成功，直接返回200
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/3/31 20:00
func Ok(c *app.RequestContext) {
	c.JSON(http.StatusOK, Response{
		errcode.SuccessCode,
		nil,
		errcode.SuccessCode.String(),
	})
}

// Fail
// @Description 请求失败，影响返回500
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/3/31 19:56
func Fail(c *app.RequestContext, code errcode.Err) {
	res := Response{
		code,
		nil,
		code.String(),
	}
	c.JSON(http.StatusInternalServerError, res)
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
		errcode.HaSysParamErr.String(),
	})
}
