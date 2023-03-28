// Path: internal/pkg/middleware
// FileName: auth.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/28$ 22:22$

package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gjing1st/hertz-admin/internal/pkg/consts"
	"net/http"
)

func LoginRequired() app.HandlerFunc {

	return func(c context.Context, ctx *app.RequestContext) {
		token := ctx.GetHeader(consts.HeaderAuthorization)
		if len(token) < len(consts.AuthPre) {
			ctx.AbortWithMsg("请携带token认证", http.StatusUnauthorized)
			return
		}
		//token验证
		//string(token[len(consts.AuthPre):])
		//var tokenService service.TokenService
		//userInfo, err := tokenService.GetInfo(token)
		//if err != 0 || userInfo.Id == 0 {
		//	//token错误或token过期，返回401
		//	ctx.AbortWithMsg("登录已失效，请重新登录", http.StatusUnauthorized)
		//	return
		//}
		//ctx.Set("userId", userInfo.Id)
		//ctx.Set("username", userInfo.Name)
		//ctx.Set("roleId", userInfo.RoleId)
		ctx.Next(c)
	}
}
