package main

import (
	"log"

	_ "gopkg.in/mgo.v2"
)

func connectMongo() {
	db, err := mgo.Dial("user:password@127.0.0.1:27017/db_learn")

	if err != nil {
		log.Fatal(err)
	}

	db.SetMode(mgo.Monotonic, true)
	defer db.Close()
}
