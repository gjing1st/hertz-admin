// Path: internal/apiserver/model/request
// FileName: category.go
// Created by bestTeam
// Author: GJing
// Date: 2022/10/30$ 20:31$

package request

type CategoryList struct {
	Name string `query:"name" form:"name"`
	PageInfo
}
