package main

import (
	"database/sql"
	"log"

	// Set Alias for MySQL Driver Package
	// To _ Since This Package Only Used in
	// For MySQL Driver
	_ "github.com/go-sql-driver/mysql"
)

func connectMySQL() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db_learn")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
