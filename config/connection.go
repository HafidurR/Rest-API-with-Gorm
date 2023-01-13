package config

import (
	"log"
	"gorm.io/gorm"
  "gorm.io/driver/mysql"
)

var db *gorm.DB
var err error

func Connect() *gorm.DB {
	db, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_store"))

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}

	// db.AutoMigrate(&Product{})
	return db
}