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

func InitApi(h *server.Hertz) {
	apiV1 := h.Group("ha/v1")
	apiV1.GET("ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, "pong")
	})

}
