// Path: internal/apiserver/model
// FileName: common.go
// Author: GJing
// Date: 2022/10/29$ 23:25$

package entity

import (
	"gorm.io/gorm"
	"time"
)

type Id struct {
	Id int `json:"id" gorm:"id"`
}

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
