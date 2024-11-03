package config

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/postgres"
)

var db *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=book-store port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
