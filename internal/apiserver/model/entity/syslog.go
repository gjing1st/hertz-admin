// Path: internal/apiserver/model/entity
// FileName: log.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 14:46$

package entity

// SysLog 操作日志表
type SysLog struct {
	BaseModel
	Username    string `json:"username" gorm:"column:username;comment:用户名;type:varchar(64);"`
	Content     string `json:"content" gorm:"column:content;comment:操作内容;type:varchar(255)"`
	RequestData string `json:"request_data" gorm:"column:request_data;comment:接口请求参数;"`
	ClientIP    string `json:"client_ip" gorm:"column:client_ip;comment:客户端ip;type:varchar(15)"`
	Result      int    `json:"result"  gorm:"column:result;default:1;comment:状态,1成功2失败;"`
	Category    int    `json:"category"  gorm:"column:category;default:1;comment:状态,1操作日志2系统日志;"`
	CheckData   string `json:"-"  gorm:"column:check_data;comment:完整性校验;type:varchar(255);"`
}

func (SysLog) TableName() string {
	return "sys_log"
}
