package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var GlobalDB *gorm.DB

func init() {

	sqliteDB, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		log.Println("failed to connect database")
		panic(err)
	}

	GlobalDB = sqliteDB

}
