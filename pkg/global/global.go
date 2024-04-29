// Path: internal/pkg/consts
// FileName: global.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/28$ 22:36$

package global

const PageSizeDefault = 10 //默认每页显示数量

const (
	DriverPostgresql = "postgresql"
	DriverOpenGauss  = "opengauss"
	DriverKingBase   = "kingbase" //人大金仓:postgres模式
	DriverMysql      = "mysql"
	DriverDM         = "dm"         //达梦数据库
	DriverSqlite     = "sqlite"     //
	DriverClickHouse = "clickhouse" //
)
