package models

import (
	"errors"
	"log"

	"github.com/mohanapranes/book-store/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

// Book represents the Book model in the database
type Book struct {
	gorm.Model
	Name         string `json:"name"`
	Author       string `json:"author"`
	Publications string `json:"publications"`
}

// init initializes the database connection and migrates the Book model
func init() {
	log.Println("Initializing database connection and migrating models...")
	config.Connect()
	db = config.GetDB()

	if err := db.AutoMigrate(&Book{}); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}
}

// CreateBook inserts a new Book record into the database
func CreateBook(book *Book) (*Book, error) {
	log.Println("Creating a new book entry...")

	if err := db.Create(book).Error; err != nil {
		log.Printf("Error creating book: %v", err)
		return nil, err
	}
	return book, nil
}

// GetAllBooks retrieves all Book records from the database
func GetAllBooks() ([]Book, error) {
	log.Println("Retrieving all books...")

	var books []Book
	if err := db.Find(&books).Error; err != nil {
		log.Printf("Error fetching books: %v", err)
		return nil, err
	}
	return books, nil
}

// GetABookById retrieves a Book record by its ID
func GetABookById(id int64) (*Book, *gorm.DB, error) {
	log.Printf("Retrieving book with ID %d...", id)

	var book Book
	if err := db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Book with ID %d not found", id)
			return nil, nil, nil
		}
		log.Printf("Error retrieving book: %v", err)
		return nil, nil, err
	}
	return &book, db, nil
}

// UpdateBook updates an existing Book record
func UpdateBook(id int64, updatedData *Book) (*Book, error) {
	log.Printf("Updating book with ID %d...", id)

	book, _, err := GetABookById(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		log.Printf("Book with ID %d not found", id)
		return nil, nil
	}

	// Update only fields with new values
	if updatedData.Name != "" {
		book.Name = updatedData.Name
	}
	if updatedData.Author != "" {
		book.Author = updatedData.Author
	}
	if updatedData.Publications != "" {
		book.Publications = updatedData.Publications
	}

	if err := db.Save(book).Error; err != nil {
		log.Printf("Error updating book: %v", err)
		return nil, err
	}
	return book, nil
}

// DeleteBookById deletes a Book record by its ID
func DeleteBookById(id int64) error {
	log.Printf("Deleting book with ID %d...", id)

	if err := db.Delete(&Book{}, id).Error; err != nil {
		log.Printf("Error deleting book: %v", err)
		return err
	}
	return nil
}
