package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Database
var Db *sql.DB

// Connecting to Database
func Connect() {
	// Getting databaase details
	var user = os.Getenv("DATABASE_USER")
	var password = os.Getenv("DATABASE_PASSWORD")
	var database = os.Getenv("DATABASE_NAME")

	// Connection string for mysql
	var connectionString = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, database)

	// Connecting to databaase
	db, err := sql.Open("mysql", connectionString)

	// Checking error
  if err != nil {
    panic(err)
  }

	Db = db
}

