package configs

import (
	"database/sql"
	"log"

	// Set Alias for MySQL Driver Package
	// To _ Since This Package Only Used in
	// For MySQL Driver
	_ "github.com/go-sql-driver/mysql"
)

// Database Configuration Struct
type databases struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Database Connect Function
func DBConnect() *sql.DB {
	db, err := sql.Open("mysql", DBConfig.User+":"+DBConfig.Password+"@tcp("+DBConfig.Host+":"+DBConfig.Port+")/"+DBConfig.Name)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
