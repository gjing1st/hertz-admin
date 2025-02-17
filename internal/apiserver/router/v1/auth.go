package v1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/gjing1st/hertz-admin/internal/pkg/middleware"
)

// initLoginRouter
// @description: 需要登录权限的路由
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:14
// @success:
func initLoginRouter(r *route.RouterGroup) {
	//需要登录权限路由组
	r.Use(middleware.LoginRequired())
	initLoginUser(r)
	initSys(r)
	//其他权限需要在具体路由组中分别注册中间件
	initAuthAdminRouter(r)  //需要管理员权限
	initSuperAdminRouter(r) //需要超级管理员权限

}

// initAuthAdminRouter
// @description: 需要管理员权限的路由
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:14
// @success:
func initAuthAdminRouter(r *route.RouterGroup) {
	r.Use(middleware.AdminRequired())
	//initSysRouter(r)
	//initUserRouter(r)
}

// initSuperAdminRouter
// @Description 需要超级管理员权限
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/6/12 14:16
func initSuperAdminRouter(r *route.RouterGroup) {
	r.Use(middleware.SuperAdminRequired())
	//initUserRouter(apiV1)
	//initRoleRouter(apiV1)

}
