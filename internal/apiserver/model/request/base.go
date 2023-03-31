// Path: internal/apiserver/model/request
// FileName: base.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/28$ 16:09$

package request

// PageInfo 公共分页基类
type PageInfo struct {
	Page     int    `query:"page" form:"page"`           // 页码
	PageSize int    `query:"page_size" form:"page_size"` // 每页大小
	Keyword  string `query:"keyword" form:"keyword"`     //关键字
	Sort     string `query:"sort" form:"sort"`           //排序参数
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}
