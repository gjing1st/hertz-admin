// Path: internal/apiserver/service
// FileName: category.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/29$ 14:25$

package service

import (
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/request"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/mysql"
)

type CategoryService struct {
}

var categoryStore mysql.CategoryStore

func (cs *CategoryService) GetList(req request.CategoryList) (list interface{}, total int64, err error) {
	list, total, err = categoryStore.List(req)

	return
}
