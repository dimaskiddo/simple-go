package configs

import (
	"log"

	// Set Alias for Mongo Driver Package
	// To mgo Since This Package Has Different Name
	// With It's Repository
	mgo "gopkg.in/mgo.v2"
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
func DBConnect() *mgo.Session {
	db, err := mgo.Dial(DBConfig.User + ":" + DBConfig.Password + "@" + DBConfig.Host + ":" + DBConfig.Port + "/" + DBConfig.Name)
	if err != nil {
		log.Fatal(err)
	}

	// Set MongoDB Parameters
	db.SetMode(mgo.Monotonic, true)

	return db
}
