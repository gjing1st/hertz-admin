// Path: internal/pkg/middleware
// FileName: auth.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/28$ 22:22$

package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/dict"
	"github.com/gjing1st/hertz-admin/internal/apiserver/service"
	"github.com/gjing1st/hertz-admin/internal/pkg/functions"
	"github.com/gjing1st/hertz-admin/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	authPre             = "Bearer "
)

// LoginRequired godoc
//
//	@Description	验证token，需要登录
//	@contact.name	GJing
//	@contact.email	gjing1st@gmail.com
//	@date			2023/3/28 22:55
func LoginRequired() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		token := ctx.GetHeader(headerAuthorization)
		if len(token) < len(authPre) {
			ctx.AbortWithMsg("请携带token认证", http.StatusUnauthorized)
			return
		}
		//token验证
		tokenStr := string(token[len(authPre):])
		var tokenService service.TokenService
		userInfo, err := tokenService.GetInfo(tokenStr)
		if err != nil || userInfo.Id == 0 {
			//token错误或token过期，返回401
			ctx.AbortWithMsg("登录已失效，请重新登录", http.StatusUnauthorized)
			return
		}
		ctx.Set("userId", userInfo.Id)
		ctx.Set("username", userInfo.Name)
		ctx.Set("roleId", userInfo.RoleId)
		c = context.WithValue(c, "testId", 123)
		ctx.Next(c)
	}
}

// AdminRequired
// @description: 需要管理员权限
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2025/2/17 10:22
func AdminRequired() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		role, b := ctx.Get("roleId")
		if !b {
			ctx.AbortWithMsg("未授权", http.StatusForbidden)
			return
		}
		roleId := utils.Int(role)
		if roleId != dict.RoleIdAdmin && roleId != dict.RoleIdSuperAdmin {
			functions.AddErrLog(log.Fields{"msg": "权限错误，需要操作员权限", "当前role_id": roleId})
			ctx.AbortWithMsg("需要管理员权限", http.StatusForbidden)
			return
		}
		ctx.Next(c)
	}
}

// SuperAdminRequired
// @description: 需要超管权限
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2025/2/17 10:21
func SuperAdminRequired() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		role, b := ctx.Get("roleId")
		if !b {
			ctx.AbortWithMsg("未授权", http.StatusForbidden)
			return
		}
		roleId := utils.Int(role)
		if roleId != dict.RoleIdSuperAdmin {
			functions.AddErrLog(log.Fields{"msg": "权限错误，需要操作员权限", "当前role_id": roleId})
			ctx.AbortWithMsg("需要超级管理员权限", http.StatusForbidden)
			return
		}
		ctx.Next(c)
	}
}
