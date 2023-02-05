package errors

import "github.com/go-sql-driver/mysql"

// Check if duplicate entry error occured
func CheckDuplicateEntryError(err error) bool {
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			// Checking if error is duplicate entry error
			if mysqlErr.Number == 1062 {
				return true
			}
		}
	}
	return false
}
