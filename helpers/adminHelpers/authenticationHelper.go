package adminHelpers

import (
	"github.com/FulgurCode/school-management-system/helpers"
	"github.com/FulgurCode/school-management-system/helpers/databaseHelpers"
	"github.com/FulgurCode/school-management-system/models"
	"github.com/gin-contrib/sessions"
)

// Storing admin id in session
func LoginWithSesssion(session sessions.Session, admin models.Admin) {
	session.Set("isLoggedIn", true)
	session.Set("id", admin.Id)

  // Saving session for 100 years
	var maxAge = 60 * 60 * 24 * 365 * 100
	session.Options(sessions.Options{MaxAge: maxAge})
	session.Save()
}

// Deleting login information from session
func DeleteAdminSession(session sessions.Session) {
	session.Options(sessions.Options{MaxAge: -1})
	session.Clear()
	session.Save()
}

// Change admin password
func ChangeAdminPassword(admin models.Admin) {
  // Getting hashed password
	var hashedPassword = helpers.HashPassword(admin.Password)
	admin.Password = hashedPassword

  // Changing password
  databaseHelpers.ChangeAdminPassword(admin)
}
