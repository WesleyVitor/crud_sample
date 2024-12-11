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
	server.GET("/books", books_controller.List)
	server.GET("/books/:id", books_controller.Show)
	server.DELETE("/books/:id", books_controller.Delete)
	server.PUT("/books/:id", books_controller.Update)

	server.Run(":8080")
}