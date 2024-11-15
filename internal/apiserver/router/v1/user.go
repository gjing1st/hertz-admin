// Path: internal/apiserver/router/v1
// FileName: user.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/31$ 19:37$

package v1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/gjing1st/hertz-admin/internal/apiserver/controller"
	"github.com/gjing1st/hertz-admin/internal/pkg/middleware"
)

var userController controller.UserController

func initUser(r *route.RouterGroup) {
	api := r.Group("user")
	api.POST("login", userController.Login)
	api.POST("logout", userController.Logout)

}
func initLoginUser(r *route.RouterGroup) {
	r.Use(middleware.LoginRequired())
	r.GET("", userController.LoginTest)
}
