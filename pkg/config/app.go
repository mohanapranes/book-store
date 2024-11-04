package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	log.Printf("Connect")
	dsn := "host=localhost user=postgres password=postgres dbname=book-store port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// d, err := gorm.Open("postgres", dsn)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
