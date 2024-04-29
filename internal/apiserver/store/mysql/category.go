// Path: internal/apiserver/store/mysql
// FileName: category.go
// Created by bestTeam
// Author: GJing
// Date: 2022/10/30$ 20:41$

package mysql

import (
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/request"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store"
)

type CategoryStore struct{}

// List
//
//	@description:	分类列表
//	@param:
//	@author:	GJing
//	@email:		gjing1st@gmail.com
//	@date:		2022/10/30 20:46
//	@success:
func (cs *CategoryStore) List(req request.CategoryList) (categories []entity.Category, total int64, err error) {
	db := store.DB.Model(&entity.Category{})
	if req.Name != "" {
		db.Where("name like ?", "%"+req.Name+"%")
	}
	//err = db.Count(&total).Error
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	err = db.Limit(limit).Offset(offset).Find(&categories).Error
	return

}

func (cs *CategoryStore) Test() (categories []entity.Category, total int64, err error) {
	db := store.DB.Model(&entity.Category{})

	//err = db.Count(&total).Error

	limit := 10
	offset := 0
	err = db.Limit(limit).Offset(offset).Find(&categories).Error
	return
}
func (cs *CategoryStore) First() (category entity.Category) {
	store.DB.First(&category)
	return
}

func (cs *CategoryStore) Index(id int) (category entity.Category) {
	store.DB.Where("id=?", id).First(&category)
	return
}
