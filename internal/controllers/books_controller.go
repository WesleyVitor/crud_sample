package controllers

import (
	"example/work-at-olist-challenge/internal/models"
	"example/work-at-olist-challenge/internal/services"

	"github.com/gin-gonic/gin"
)

type BooksContoller struct {
	bookService services.BookService
	authorService services.AuthorService
}

func NewBooksController(bookService services.BookService, authorService services.AuthorService) BooksContoller {
	return BooksContoller{
		bookService: bookService,
		authorService: authorService,
	}
}

func (bc BooksContoller) Create(c *gin.Context) {

	var body struct{
		Name string
		Edition int
		PublicationYear int
		Authors []int
	}
	err := c.BindJSON(&body)
	
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	authors, err := bc.authorService.GetAuthorsByIDs(body.Authors)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	book := models.Book{
		Name: body.Name,
		Edition: body.Edition,
		PublicationYear: body.PublicationYear,
		Authors: authors,
	}
	book_inserted := bc.bookService.CreateBook(book)
	c.JSON(201, book_inserted)
}