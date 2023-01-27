package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FulgurCode/school-management-system/database"
	"github.com/FulgurCode/school-management-system/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load enviornment variables
	godotenv.Load(".env")

	// Connecting to database
	database.Connect()

	// Setting up gin router
	gin.SetMode(gin.DebugMode)
	var router = router.Router()

	// cors setup
	var c = cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	})

	// Listen to requests in 'port'
	var port = os.Getenv("PORT")
	if err := http.ListenAndServe(port, c.Handler(router)); err != nil {
		log.Fatal(err)
	}

	// Closing database connection
	defer database.Db.Close()
}
