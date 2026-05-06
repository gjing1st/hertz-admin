// Path: internal/apiserver/model/response
// FileName: config.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 15:09$

package response

type Init struct {
	Initialized bool `json:"initialized"`
}

// SysRunDate 运行时长
type SysRunDate struct {
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type LoginTypeRes struct {
	LoginType string `json:"login_type"`
}

// InitStepValue 初始化步骤对应的值
type InitStepValue struct {
	User    int `json:"user"`
	Network int `json:"network"`
}
