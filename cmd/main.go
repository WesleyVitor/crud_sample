package main

import (
	"example/work-at-olist-challenge/controllers"
	"example/work-at-olist-challenge/pkg/db"
	"example/work-at-olist-challenge/pkg/services"

	"github.com/gin-gonic/gin"
)

func main() {
	
	db := db.Connect()
	
	book_service := services.NewBookService(db)
	books_controller := controllers.NewBooksController(book_service)



	server := gin.Default()
	server.GET("/books", books_controller.Create)

	server.Run(":8080")
}