// Path: internal/apiserver/router/v1
// FileName: category.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/30$ 20:26$

package v1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/gjing1st/hertz-admin/internal/apiserver/controller"
)

func initCategory(r *route.RouterGroup) {
	api := r.Group("category")
	categoryController := controller.CategoryController{}
	api.GET("first", categoryController.First)  //压测测试使用
	api.GET("index", categoryController.Index)  //压测测试使用
	api.GET("list", categoryController.GetList) //压测测试使用

}
