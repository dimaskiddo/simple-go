package helpers

import (
	"log"

	// Set Alias for Mongo Driver Package
	// To mgo Since This Package Has Different Name
	// With It's Repository
	mgo "gopkg.in/mgo.v2"
)

// Database Configuration Struct
type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Database Configuration Variable
var DBConfig Connection

// Database Connect Function
func MongoDBConnect() (*mgo.Session, *mgo.Database) {
	sess, err := mgo.Dial(DBConfig.User + ":" + DBConfig.Password + "@" + DBConfig.Host + ":" + DBConfig.Port + "/" + DBConfig.Name)
	if err != nil {
		log.Fatal(err)
	}

	// Set MongoDB Parameters
	sess.SetMode(mgo.Monotonic, true)

	return sess, sess.DB(DBConfig.Name)
}
