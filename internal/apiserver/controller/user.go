// Path: internal/apiserver/controller
// FileName: user.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/31$ 19:38$

package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/request"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/response"
	"github.com/gjing1st/hertz-admin/internal/apiserver/service"
	"github.com/gjing1st/hertz-admin/pkg/errcode"
	"strings"
)

type UserController struct {
}

var userService = service.UserService{}

// Login godoc
// @Summary 用户登录
// @Description 用户登录
// @Param data body request.Login true "用户名和密码"
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} UserController "操作成功"
// @Failure 500 {object} string
// @Router /user/login [post]
func (uc *UserController) Login(ctx context.Context, c *app.RequestContext) {
	var req request.UserLogin
	if err := c.Bind(&req); err != nil {
		response.ParamErr(c)
		return
	}
	res, err := userService.Login(&req)
	content := "登录"
	if err != nil {
		if errors.Is(err, errcode.UserNotFound) {
			response.FailWithLog(err, content, nil, c)
			return
		}
		//response.FailWithLog(errCode, global.LoginFail, req.Name, content, nil, c)
		response.FailWithDataLog(res, err, content, nil, c)
	} else {
		response.OkWithDataLog(res, content, nil, c)
	}
}

func (uc *UserController) LoginTest(ctx context.Context, c *app.RequestContext) {
	testId := ctx.Value("testId")
	fmt.Println("testId====", testId)
}

// Logout godoc
// @Summary 登出
// @Description
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string
// @Router /user/logout [post]
func (uc *UserController) Logout(ctx context.Context, c *app.RequestContext) {
	auth := c.GetHeader("Authorization")
	token := strings.Replace(string(auth), "Bearer ", "", 1)
	service.TokenService{}.RemoveToken(token)
	content := "登出"
	response.OkWithLog("", content, c)

}

func (uc *UserController) Register(ctx context.Context, c *app.RequestContext) {
	var req request.UserRegister
	if err := c.BindJSON(&req); err != nil {
		response.ParamErr(c)
		return
	}

	err := userService.Register(&req)
	ctx = context.WithValue(ctx, "content", "用户注册")
	ctx = context.WithValue(ctx, "req", req)
	ctx = context.WithValue(ctx, "err", err)
	if err != nil {
		response.FailWithLog(err, "用户注册", req, c)
		return
	}
	response.OkWithDataLog("", "用户注册", req, c)
}
