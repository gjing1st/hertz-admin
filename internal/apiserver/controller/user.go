// Path: internal/apiserver/controller
// FileName: user.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/31$ 19:38$

package controller

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/request"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/response"
)

type UserController struct {
}

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
	var p request.Login
	if err := c.Bind(&p); err != nil {
		response.ParamErr(c)
		return
	}
	response.Ok(c)
}

func (uc *UserController) LoginTest(ctx context.Context, c *app.RequestContext) {
	testId := ctx.Value("testId")
	fmt.Println("testId====", testId)
}
