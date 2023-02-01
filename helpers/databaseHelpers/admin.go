package databaseHelpers

import (
	"fmt"

	"github.com/FulgurCode/school-management-system/connections"
	"github.com/FulgurCode/school-management-system/models"
)

// Getting admin user using username
func GetAdminUser(username string) models.Admin {
  fmt.Println(username)
	// database
	var db = connections.Db

	// Query for database
	var query = "SELECT * FROM admin WHERE username = ?;"

	// Getting admin user
	var admin models.Admin
	db.QueryRow(query, username).Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Email, &admin.Phone)
	return admin
}
