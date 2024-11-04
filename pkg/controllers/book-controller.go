package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohanapranes/book-store/pkg/models"
	"github.com/mohanapranes/book-store/pkg/utils"
)

var NewBook models.Book

// CreateBook handles the creation of a new book entry
func CreateBook(cntx *gin.Context) {
	log.Println("Creating a new book")
	bookRequest := &models.Book{}

	// Parse the JSON request body
	if err := utils.ParseBody(cntx.Request, bookRequest); err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call model to create a new book
	createdBook, _ := models.CreateBook(bookRequest)
	cntx.JSON(http.StatusCreated, createdBook)
}

// GetBooks retrieves all books from the database
func GetBooks(cntx *gin.Context) {
	log.Println("Fetching all books")
	books, _ := models.GetAllBooks()
	cntx.JSON(http.StatusOK, books)
}

// GetABookById retrieves a single book by its ID
func GetABookById(cntx *gin.Context) {
	log.Println("Fetching book by ID")
	bookId, err := parseBookId(cntx)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the book from the model
	book, _, err := models.GetABookById(bookId)
	if err != nil {
		cntx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	cntx.JSON(http.StatusOK, book)
}

// UpdateABook updates an existing book by its ID
func UpdateABook(cntx *gin.Context) {
	log.Println("Updating a book")
	bookId, err := parseBookId(cntx)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateBookReq := &models.Book{}
	if err := utils.ParseBody(cntx.Request, updateBookReq); err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Retrieve the existing book from the model
	book, db, err := models.GetABookById(bookId)
	if book == nil {
		cntx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Update only non-empty fields
	if updateBookReq.Name != "" {
		book.Name = updateBookReq.Name
	}
	if updateBookReq.Author != "" {
		book.Author = updateBookReq.Author
	}
	if updateBookReq.Publications != "" {
		book.Publications = updateBookReq.Publications
	}

	db.Save(&book)
	cntx.JSON(http.StatusOK, book)
}

// DeleteABook deletes a book by its ID
func DeleteABook(cntx *gin.Context) {
	log.Println("Deleting a book")
	bookId, err := parseBookId(cntx)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the model to delete the book
	err = models.DeleteBookById(bookId)
	if err != nil {
		cntx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// parseBookId extracts and validates the book ID from the URL parameter
func parseBookId(cntx *gin.Context) (int64, error) {
	bookIdParam := cntx.Param("bookId")
	bookId, err := strconv.ParseInt(bookIdParam, 10, 64)
	if err != nil {
		return 0, err
	}
	return bookId, nil
}
