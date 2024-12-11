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
		c.JSON(400, err.Error())
	}
	authors, err := bc.authorService.GetAuthorsByIDs(body.Authors)
	if err != nil {
		c.JSON(400, err.Error())
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

func (bc BooksContoller) List(c *gin.Context) {
	books := bc.bookService.GetAllBooks()
	c.JSON(200, books)

}

func (bc BooksContoller) Show(c *gin.Context) {
	book_id := c.Param("id")

	book, err := bc.bookService.GetBookByID(book_id)

	if err != nil {
		c.JSON(404, err.Error())
	}
	c.JSON(200, book)	
	
}

func (bc BooksContoller) Delete(c *gin.Context){
	book_id := c.Param("id")

	err := bc.bookService.DeleteBookByID(book_id)

	if err != nil {
		c.JSON(404, err.Error())
	}
	c.JSON(204, nil)
}

func (bc BooksContoller) Update(c *gin.Context){
	var body struct{
		Name string
		Edition int
		PublicationYear int
		Authors []int
	}
	err := c.BindJSON(&body)
	
	if err != nil {
		c.JSON(400, err.Error())
	}
	var authors []*models.Author
	if len(body.Authors) != 0 {
		authors, err = bc.authorService.GetAuthorsByIDs(body.Authors)
		if err != nil {
			c.JSON(400, err.Error())
		}
	} else {
		authors = []*models.Author{}
	}


	book := models.Book{
		Name: body.Name,
		Edition: body.Edition,
		PublicationYear: body.PublicationYear,
		Authors: authors,
	}
	book_id := c.Param("id")
	err = bc.bookService.UpdateBookByID(book_id, book)
	if err != nil {
		c.JSON(404, err.Error())
	}
	c.JSON(200, book)
}