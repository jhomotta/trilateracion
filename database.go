package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(ConnectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Can't connect to DB!")
		log.Fatal(err.Error())
	} else {
		log.Println("Connect to DB!")
	}

	DB.AutoMigrate(&Satellite{})
}
