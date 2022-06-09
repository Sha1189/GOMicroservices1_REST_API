package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

// Connect to database
func ConnectDB() {

	db, err := sql.Open("mysql", "root:password@tcp(db-gomicrosvs1)/db_courses")

	// handle error
	if err != nil {
		panic(err.Error())

	}
	dB = db
}

// return database variable to link db to APIs
func GetDB() *sql.DB {

	return dB
}
