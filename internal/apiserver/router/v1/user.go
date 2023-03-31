// Path: internal/apiserver/router/v1
// FileName: user.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/31$ 19:37$

package v1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/gjing1st/hertz-admin/internal/apiserver/controller"
)

func initUser(r *route.RouterGroup) {
	api := r.Group("user")
	var userController controller.UserController
	api.POST("login", userController.Login)

}
