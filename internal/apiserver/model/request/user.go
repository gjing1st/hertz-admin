// Path: internal/apiserver/model/request
// FileName: user.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/31$ 19:46$

package request

// Login 登录请求参数
type Login struct {
	Username string `json:"username,required" `
	Password string `json:"password,required"`
}
