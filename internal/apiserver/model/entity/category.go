// $
// category.go
// Created by dkedTeam.
// Author: GJing
// Date: 2022/10/28$ 11:28$

package entity

// Category 结构体
type Category struct {
	BaseModel
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:55;"`
	Status   int    `json:"status" form:"status" gorm:"column:status;comment:状态;size:1;"`
	ParentId int    `json:"parent_id" form:"parent_id" gorm:"column:parent_id;comment:父id;size:16;"`
}

// TableName Category 表名
func (Category) TableName() string {
	return "category"
}
