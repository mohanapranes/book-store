package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohanapranes/book-store/pkg/controllers"
)

var RegisterBookStoreRoutes = func() {
	router := gin.Default()

	router.POST("/book", controllers.CreateBook)
	router.GET("/book", controllers.GetBooks)
	router.GET("/book/{bookId}", controllers.GetABookById)
	router.PUT("/book/{bookId}", controllers.UpdateABook)
	router.DELETE("/book/{bookId}", controllers.DeleteABook)
}
