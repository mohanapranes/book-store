package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohanapranes/book-store/pkg/controllers"
)

// RegisterRoutes sets up API endpoints for book-related operations
func RegisterRoutes(router *gin.Engine) {
	// Group book-related routes under /books
	bookRoutes := router.Group("/books")
	{
		bookRoutes.POST("", controllers.CreateBook)            // POST /books
		bookRoutes.GET("", controllers.GetBooks)               // GET /books
		bookRoutes.GET("/:bookId", controllers.GetABookById)   // GET /books/:bookId
		bookRoutes.PUT("/:bookId", controllers.UpdateABook)    // PUT /books/:bookId
		bookRoutes.DELETE("/:bookId", controllers.DeleteABook) // DELETE /books/:bookId
	}
}
