package main

import (
	"example/work-at-olist-challenge/internal/controllers"
	"example/work-at-olist-challenge/internal/db"
	"example/work-at-olist-challenge/internal/services"

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