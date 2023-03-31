// Path: internal/apiserver/router
// FileName: init.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/28$ 21:58$

package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// InitApi
// @Description 初始化路由
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/3/28 22:55
func InitApi(h *server.Hertz) {
	apiV1 := h.Group("ha/v1")
	apiV1.GET("ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, "pong")
	})

	{
		initCategory(apiV1)
		initUser(apiV1)
	}

}
