package models

import (
	"log"

	"github.com/mohanapranes/book-store/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name         string `gorm:"" json:"name"`
	Author       string `json:"author"`
	Publications string `json:"publications"`
}

func init() {
	log.Printf("Init")
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(book *Book) *Book {
	log.Printf("CreateBook")
	db.Create(book)
	return book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetABookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("id=?", id).Find(&book)
	return &book, db
}

func DeleteBookById(id int64) *Book {
	var book Book
	db.Where("id=?", id).Delete(&book)
	return &book
}
