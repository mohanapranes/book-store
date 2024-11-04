package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohanapranes/book-store/pkg/routes"
)

func main() {
	router := gin.Default()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	http.ListenAndServe("localhost:9010", router)
}
