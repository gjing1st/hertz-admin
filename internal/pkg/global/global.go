// Path: internal/pkg/consts
// FileName: global.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/28$ 22:36$

package global

const (
	HeaderAuthorization = "Authorization"
	AuthPre             = "Bearer "
)
const (
	TimeFormat = "2006-01-02 15:04:05"
	MaxLimit   = 1000
)

const PageSizeDefault = 10 //默认每页显示数量

const (
	SecondsPerMinute = 60
	SecondsPerHour   = 60 * SecondsPerMinute
	SecondsPerDay    = 24 * SecondsPerHour
	SecondsPerWeek   = 7 * SecondsPerDay
	DaysPer400Years  = 365*400 + 97
	DaysPer100Years  = 365*100 + 24
	DaysPer4Years    = 365*4 + 1
)
const SuperAdmin = "superAdmin1"

const (
	DriverPostgresql = "postgresql"
	DriverOpenGauss  = "opengauss"
	DriverKingBase   = "kingbase"
	DriverMysql      = "mysql"
	DriverMongo      = "mongodb"
	DriverDM         = "dm"         //达梦数据库
	DriverSqlite     = "sqlite"     //
	DriverClickHouse = "clickhouse" //
)
