// Path: internal/apiserver/model/entity
// FileName: user.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/26$ 19:35$

package entity

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// User 用户表
type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name" gorm:"column:name;comment:用户名;type:varchar(255);NOT NULL;"`
	//Name          string                `json:"name" gorm:"column:name;comment:用户名;uniqueIndex:user_name;type:varchar(64);NOT NULL;"`
	NickName  string `json:"nick_name" gorm:"column:nick_name;comment:昵称;type:varchar(255);NOT NULL;"`
	RoleId    int    `json:"role_id" gorm:"column:role_id;comment:角色类型;" `
	Password  string `json:"-"  gorm:"column:password;comment:密码;type:varchar(255);"`
	Token     string `json:"token"  gorm:"column:token;comment:令牌;type:varchar(255);"`
	Status    int    `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;"`
	LoginType int    `json:"login_type" form:"login_type" gorm:"column:login_type;default:1;comment:登录方式,1口令2前端key3后端key;"`
	//LoginType     int                   `json:"login_type" form:"login_type" gorm:"uniqueIndex:user_name;column:login_type;default:1;comment:登录方式,1口令2前端key3后端key;"`
	ErrNum       int                   `json:"-"  gorm:"column:err_num;default:0;comment:错误次数，向上累加;"`
	PwdUpdatedAt time.Time             `json:"-"  gorm:"column:pwd_updated_at;comment:密码更新时间;"`
	CheckData    string                `json:"-"  gorm:"column:check_data;comment:完整性校验;"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at,omitzero"`
	DeletedAt    soft_delete.DeletedAt `json:"-"`
	//DeletedAt     soft_delete.DeletedAt `gorm:"uniqueIndex:user_name" json:"-"`
}

func (User) TableName() string {
	return "user"
}

type UserTokenInfo struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	RoleId    int    `json:"role_id"`
	LoginType int    `json:"login_type"`
}
