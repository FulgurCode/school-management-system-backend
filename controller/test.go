package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTestApi(c *gin.Context) {
	fmt.Println("Response received")
	c.JSON(200, "Respose from server")
}
