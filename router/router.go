package router

import (
	"github.com/FulgurCode/school-management-system/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// New router declared
	var router = gin.Default()

	// Creating sessions
	var store = cookie.NewStore([]byte("school"))
	var sessionNames = []string{"admin"}
	router.Use(sessions.SessionsMany(sessionNames, store))

	// Forwarding request on '/api/test'
	router.GET("/api/test", controller.GetTestApi)

	// // Forwarding request on '/api/admin'
	router.POST("/api/admin/login", controller.AdminLoginRoute)
  router.GET("/api/admin/logout", controller.AdminLogoutRoute)
	router.GET("/api/admin/checklogin", controller.AdminCheckLoginRoute)

	return router
}
