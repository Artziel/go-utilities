package GoUtilities

import (
	go_mysql "github.com/go-sql-driver/mysql"
)

func MySQLErrorCode(err error) int {
	result := 0
	if mysqlError, ok := err.(*go_mysql.MySQLError); ok {
		result = int(mysqlError.Number)
	}

	return result
}
