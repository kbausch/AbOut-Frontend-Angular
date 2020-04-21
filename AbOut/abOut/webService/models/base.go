package models

import (
	"fmt"
	"os"

	"database/sql"

	// This import makes sure gorm uses the right dialect of sql when building
	// queiries.

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// Init sets up the database connection.
func init() {
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			fmt.Println(err)
		}
	}

	// Get the db connection URL from the environment.
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbHost := os.Getenv("db_host")
	dbName := os.Getenv("db_name")
	dbType := os.Getenv("db_type")
	dbPort := os.Getenv("db_port")

	dbURI := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8",
		username, password, dbHost, dbPort, dbName)
	fmt.Println(dbURI)

	conn, err := sql.Open(dbType, dbURI)
	if err != nil {
		fmt.Println(err)
	}
	db = conn

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

// GetDB returns a handle to the DB object.
func GetDB() *sql.DB {
	return db
}
