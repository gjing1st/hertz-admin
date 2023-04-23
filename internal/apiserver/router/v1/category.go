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
	//服务器40核64GB 部署单个副本
	// ping接口 35w+ tps
	api.GET("first", categoryController.First)   //压测测试使用 6w tps
	api.GET("index", categoryController.Index)   //压测测试使用 6w tps
	api.GET("list", categoryController.GetList)  //压测测试使用 表7条数据，5w tps
	api.GET("cpu", categoryController.Calculate) //压测200数求和测试使用
	api.GET("cache", categoryController.Cache)   //压测gcache缓存使用   19w tps

}
