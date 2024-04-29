// Path: internal/apiserver/model/dict
// FileName: adminlog.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 16:19$

package dict

const (
	AdminLogResultOk   = 1 //操作成功
	SysLogResultOk     = 1 //操作成功
	AdminLogResultFail = 2 //操作失败
)

const (
	SysLogCategoryOperation = 1 //操作日志
	SysLogCategorySys       = 2 //系统日志
)
