// Path: internal/apiserver/model/entity
// FileName: config.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 9:35$

package entity

// Config 配置字典
type Config struct {
	BaseModel
	//Name   string      `json:"name" gorm:"column:name;comment:key;uniqueIndex:key_name;type:varchar(64);NOT NULL;"`
	Name   string      `json:"name" gorm:"column:name;comment:key;type:varchar(64);NOT NULL;"`
	Value  interface{} `json:"value" gorm:"column:value;comment:value;type:varchar(255);"`
	Status int         `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;"`
	Remark int         `json:"remark" form:"remark" gorm:"column:status;default:1;comment:状态;"`
}

func (Config) TableName() string {
	return "config"
}
