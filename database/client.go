package database

import (
	"fmt"
	"golang-crud-rest-api/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(ConnectionString string) {
	fmt.Println("I am here", ConnectionString)
	Instance, err = gorm.Open(mysql.Open(ConnectionString),
		&gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	log.Println("Database Migration Completed...")
}
