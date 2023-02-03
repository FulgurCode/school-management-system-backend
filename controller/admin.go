package controller

import (
	"github.com/FulgurCode/school-management-system/helpers"
	"github.com/FulgurCode/school-management-system/helpers/adminHelpers"
	"github.com/FulgurCode/school-management-system/helpers/databaseHelpers"
	"github.com/FulgurCode/school-management-system/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// POST request on /api/admin/login
func AdminLoginRoute(c *gin.Context) {
	// Getting request body
	var data models.Admin
	if err := c.BindJSON(&data); err != nil {
		panic(err)
	}

	// Getting password and comparing password
	var admin = databaseHelpers.GetAdminUser(data.Username)
	var result = helpers.ComparePassword(admin.Password, data.Password)
	if !result {
		c.JSON(401, "Wrong credentials")
		return
	}

	// storing id and sending response if password is correct
	var session = sessions.DefaultMany(c, "admin")
	adminHelpers.LoginWithSesssion(session, admin)
	c.JSON(200, "Login Successful")
}

// GET request on '/api/admin/logout'
func AdminLogoutRoute(c *gin.Context) {
	// clearing 'admin' session
	var session = sessions.DefaultMany(c, "admin")
	adminHelpers.DeleteAdminSession(session)

	// Response for the request
	c.JSON(200, "Loggged Out")
}

// GET request on '/api/admin/checklogin'
func AdminCheckLoginRoute(c *gin.Context) {
	// Getting session value
	var session = sessions.DefaultMany(c, "admin")
	if session.Get("isLoggedIn") == true {
		c.JSON(200, "Logged In")
	} else {
		c.JSON(401, "Not Logged In")
	}
}

// POST request on '/api/admin/change-password'
func ChangeAdminPassword(c *gin.Context) {
	// Getting request body
	var data models.Admin
	if err := c.BindJSON(&data); err != nil {
		panic(err)
	}

	// Checking if loggedin
	var session = sessions.DefaultMany(c, "admin")
	var isLoggedIn = session.Get("isLoggedIn")
	if isLoggedIn != true {
		c.JSON(401, "Not Logged In")
		return
	}

	// Change admin Password and sending response
  data.Id = session.Get("id").(string)
	adminHelpers.ChangeAdminPassword(data)
	c.JSON(200, "Admin password changed successfully")
}
