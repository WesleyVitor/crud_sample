package controllers

import (
	"example/work-at-olist-challenge/pkg/models"
	"example/work-at-olist-challenge/pkg/services"

	"github.com/gin-gonic/gin"
)

type BooksContoller struct {
	bookService services.BookService
}

func NewBooksController(bookService services.BookService) BooksContoller {
	return BooksContoller{
		bookService: bookService,
	}
}

func (bc BooksContoller) Create(c *gin.Context) {
	book_inserted := bc.bookService.CreateBook(models.Book{Name: "Em busca do tempo perdido", Edition: 1, PublicationYear: 1954})
	c.JSON(201, book_inserted)
}