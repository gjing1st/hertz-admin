// Path: internal/pkg/middleware
// FileName: auth.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/28$ 22:22$

package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
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
		c = context.WithValue(c, "testId", 123)
		ctx.Next(c)
	}
}
