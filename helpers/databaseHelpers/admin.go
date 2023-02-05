package databaseHelpers

import (
	"github.com/FulgurCode/school-management-system/connections"
	"github.com/FulgurCode/school-management-system/models"
)

// Getting admin user using username
func GetAdminUser(username string) models.Admin {
	// database
	var db = connections.Db

	// Query for database
	var query = "SELECT * FROM admin WHERE username = ?;"

	// Getting admin user
	var admin models.Admin
	db.QueryRow(query, username).Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Email, &admin.Phone)
	return admin
}

// Chnage admin password using admin id
func ChangeAdminPassword(admin models.Admin) {
	// database
	var db = connections.Db

	// creating query for database and runnning it
	var query = "UPDATE admin SET password = ? WHERE id = ?;"
	db.Exec(query, admin.Password, admin.Id)
}
