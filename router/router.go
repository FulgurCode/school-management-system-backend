package router

import (
	"github.com/FulgurCode/school-management-system/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// New router declared
	var router = gin.Default()

	// Forwarding request on '/api/test'
	router.GET("/api/test", controller.GetTestApi)

	return router
}
