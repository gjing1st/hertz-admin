// Path: internal/apiserver/model/request
// FileName: adminlog.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 16:11$

package request

import "time"

// SysLogCreate 添加管理员日志
type SysLogCreate struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	ClientIP string `json:"client_ip"`
	Result   int    `json:"result"`
}

// SysLogList 日志列表
type SysLogList struct {
	PageInfo
	Category  int       `json:"category" form:"category"`
	StartDate time.Time `json:"start_time" form:"start_time"`
	EndDate   time.Time `json:"end_time" form:"end_time"`
}

type SysLogExport struct {
	Keyword   string    `json:"keyword" form:"keyword"` //关键字
	Category  int       `json:"category" form:"category"`
	StartDate time.Time `json:"start_time" form:"start_time"`
	EndDate   time.Time `json:"end_time" form:"end_time"`
}
