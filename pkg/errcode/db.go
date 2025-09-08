package errcode

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// IsDuplicateKeyError 判断错误是否为主键冲突错误
func IsDuplicateKeyError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		// MySQL 错误代码 1062 表示主键或唯一键冲突
		return mysqlErr.Number == 1062
	}
	return false
}
