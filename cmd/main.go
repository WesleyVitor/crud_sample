package main

import (
	"example/work-at-olist-challenge/pkg/controllers"
	"example/work-at-olist-challenge/pkg/db"
	"example/work-at-olist-challenge/pkg/services"

	"github.com/gin-gonic/gin"
)

func main() {
	
	db := db.Connect()
	
	book_service := services.NewBookService(db)
	author_service := services.NewAuthorService(db)
	books_controller := controllers.NewBooksController(book_service, author_service)



	server := gin.Default()
	server.POST("/books", books_controller.Create)

	server.Run(":8080")
}