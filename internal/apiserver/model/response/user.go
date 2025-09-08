// Path: internal/apiserver/model/response
// FileName: user.go
// Created by bestTeam
// Author: GJing
// Date: 2023/6/6$ 16:34$

package response

import "github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"

type UserLogin struct {
	RetryCount int `json:"retry_count"`
	*entity.User
	PwdDue bool `json:"pwd_due"`
}

type ChangePwd struct {
	RetryCount int `json:"retry_count"`
}

// KeyQuery 查询ukey是否存在
type KeyQuery struct {
	Exist bool `json:"exist"`
	Bound bool `json:"bound"`
}
