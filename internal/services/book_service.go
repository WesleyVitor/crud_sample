package services

import (
	"errors"
	"example/work-at-olist-challenge/internal/models"

	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

// Create a new BookService instance
func NewBookService(db *gorm.DB) BookService {
	return BookService{
		db: db,
	}
}

// Create a new book
func (bs BookService) CreateBook(book models.Book) models.Book {
	bs.db.Create(&book)

	return book
}

// Get all books and preloads authors
func (bs BookService) GetAllBooks() []models.Book {
	var books []models.Book
	bs.db.Preload("Authors").Find(&books)

	return books

}

// Get a book by ID and preloads authors and return an error if the book is not found
func (bs BookService) GetBookByID(id string) (models.Book, error) {

	var book models.Book
	result := bs.db.Preload("Authors").First(&book, id)
	if result.Error != nil {
		return book, errors.New("book not found")
	}

	return book, nil
}

// Delete a book by ID and return an error if the book is not found
func (bs BookService) DeleteBookByID(id string) error {
	var book models.Book
	result := bs.db.First(&book, id)
	if result.Error != nil {
		return errors.New("book not found")
	}
	bs.db.Delete(&book)

	return nil
}

// Update a book by ID and return an error if the book is not found
func (bs BookService) UpdateBookByID(id string, book models.Book) error {
	var bookToUpdate models.Book
	result := bs.db.First(&bookToUpdate, id)
	if result.Error != nil {
		return errors.New("book not found")
	}
	bs.db.Model(&bookToUpdate).Updates(book)

	return nil
}

