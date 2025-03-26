// Path: internal/apiserver/model/request
// FileName: user.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/31$ 19:46$

package request

// Login 登录请求参数
type Login struct {
	Username string `json:"username,required" `
	Password string `json:"password,required"`
}

// UserCreate 创建用户
type UserCreate struct {
	Name      string `json:"name" binding:"required"`
	Serial    string `json:"serial"` //UKey序列号
	SignData  string `json:"sign"`
	TimeStamp string `json:"timestamp"`
	RoleId    int    `json:"role_id" form:"role_id"`
	Cert      string `json:"cert"`
	Pin       string `json:"pin"`
	Token     string `json:"token"`
	LoginType int    `json:"login_type"` //登录方式
}

// UserLogin 用户登录请求
type UserLogin struct {
	Name      string `json:"name" binding:"required"`
	LoginType int    `json:"login_type"`
	Pin       string `json:"pin"`
	Password  string `json:"pwd"`
}

// UserList 用户列表
type UserList struct {
	PageInfo
	RoleId []int `json:"role_id" form:"role_id"`
}

type UKeyLogin struct {
	Name   string `json:"name" binding:"required"`
	Serial string `json:"serial" binding:"required"`
	//SignData  string `json:"sign" binding:"required"`
	//TimeStamp string `json:"timestamp" binding:"required"`
	Token     string `json:"token" binding:"required"`
	Pin       string `json:"pin"`
	LoginType int    `json:"login_type"`
}

// CreateUKeyUser 初始化时的添加后端ukey管理员
type CreateUKeyUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"pin"  binding:"required"`
}

// ChangePasswd 修改密码
type ChangePasswd struct {
	OldPassword string `json:"old_pin"  binding:"required"`
	NewPassword string `json:"pin"  binding:"required"`
	Type        int    `json:"type"` //pin或者密码.1修改密码，2修改pin
}

// KeyBackup 密钥备份与恢复
type KeyBackup struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"pin"  binding:"required"`
}

// UserDelete 删除管理员
type UserDelete struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	KeySn    string `json:"keysn"`
}

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}
