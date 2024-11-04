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

func CreateBook(cntx *gin.Context) {
	log.Printf("Create Book")
	bookRequest := &models.Book{}
	utils.ParseBody(cntx.Request, bookRequest)
	models.CreateBook(bookRequest)
	cntx.IndentedJSON(http.StatusCreated, bookRequest)

}

func GetBooks(cntx *gin.Context) {

	books := models.GetAllBooks()
	cntx.IndentedJSON(http.StatusOK, books)

}

func GetABookById(cntx *gin.Context) {
	log.Printf("GetABookById")
	bookIdParam := cntx.Param("bookId")
	bookId, _ := strconv.ParseInt(bookIdParam, 0, 0)
	book, _ := models.GetABookById(bookId)
	cntx.IndentedJSON(http.StatusOK, book)

}

func UpdateABook(cntx *gin.Context) {

	updateBookReq := &models.Book{}
	utils.ParseBody(cntx.Request, updateBookReq)

	bookIdParam := cntx.Param("bookId")
	bookId, _ := strconv.ParseInt(bookIdParam, 0, 0)
	book, db := models.GetABookById(bookId)

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
	cntx.IndentedJSON(http.StatusCreated, book)

}

func DeleteABook(cntx *gin.Context) {

	bookIdParam := cntx.Param("bookId")
	bookId, _ := strconv.ParseInt(bookIdParam, 0, 0)
	book := models.DeleteBookById(bookId)
	cntx.IndentedJSON(http.StatusOK, book)

}
