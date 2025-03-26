// Path: internal/apiserver/router/v1
// FileName: user.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/31$ 19:37$

package v1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/gjing1st/hertz-admin/internal/apiserver/controller"
)

var userController controller.UserController

func initUser(r *route.RouterGroup) {
	api := r.Group("user")
	api.POST("login", userController.Login)
	r.POST("logout", userController.Logout)

}
func initLoginUser(r *route.RouterGroup) {
	r.GET("", userController.LoginTest)
}
