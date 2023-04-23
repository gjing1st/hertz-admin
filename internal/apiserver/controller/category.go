// Path: internal/apiserver/controller
// FileName: category.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/30$ 20:26$

package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/request"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/response"
	"github.com/gjing1st/hertz-admin/internal/apiserver/service"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/cache"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/mysql"
	"github.com/gjing1st/hertz-admin/pkg/utils"
)

type CategoryController struct {
}

// First 测试handler
//
//	@Summary		测试Summary
//	@Description	测试Description
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/category/first [get]
func (cc *CategoryController) First(ctx context.Context, c *app.RequestContext) {
	var cate mysql.CategoryStore
	res := cate.First()
	response.OkWithData(c, res)
}

// Index 测试handler
//
//	@Summary		测试index
//
// @Param id query int true "id主键"
//
//	@Description	测试Description
//	@Router			/category/index [get]
func (cc *CategoryController) Index(ctx context.Context, c *app.RequestContext) {
	var cate mysql.CategoryStore
	ids := c.Query("id")

	res := cate.Index(utils.Int(ids))
	response.OkWithData(c, res)

}

var categoryService service.CategoryService

// GetList 测试handler
//
//	@Summary		测试index
//
// @ param  id int
//
//	@Description	测试Description
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/ping [get]
func (cc *CategoryController) GetList(ctx context.Context, c *app.RequestContext) {
	var req request.CategoryList
	_ = c.Bind(&req)
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	list, total, err := categoryService.GetList(req)
	if err != nil {

	} else {
		response.OkWithData(c, response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		})
	}

}

func (cc *CategoryController) Calculate(ctx context.Context, c *app.RequestContext) {
	sum := 1
	for i := 0; i < 500000; i++ {
		sum += i
	}
	response.OkWithData(c, sum)
}

type Cache struct {
}

func (cc *CategoryController) Cache(ctx context.Context, c *app.RequestContext) {
	a, err := cache.GetCache().Get("a")
	if err != nil {
		cache.GetCache().Set("a", "123123")
	}
	response.OkWithData(c, a)

}
