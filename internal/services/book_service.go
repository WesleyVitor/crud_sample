package services

import (
	"errors"
	"example/work-at-olist-challenge/internal/models"

	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) BookService {
	return BookService{
		db: db,
	}
}

func (bs BookService) CreateBook(book models.Book) models.Book {
	bs.db.Create(&book)

	return book
}

func (bs BookService) GetAllBooks() []models.Book {
	var books []models.Book
	bs.db.Preload("Authors").Find(&books)

	return books

}

func (bs BookService) GetBookByID(id string) (models.Book, error) {

	var book models.Book
	result := bs.db.Preload("Authors").First(&book, id)
	if result.Error != nil {
		return book, errors.New("book not found")
	}

	return book, nil
}

func (bs BookService) DeleteBookByID(id string) error {
	var book models.Book
	result := bs.db.First(&book, id)
	if result.Error != nil {
		return errors.New("book not found")
	}
	bs.db.Delete(&book)

	return nil
}

func (bs BookService) UpdateBookByID(id string, book models.Book) error {
	var bookToUpdate models.Book
	result := bs.db.First(&bookToUpdate, id)
	if result.Error != nil {
		return errors.New("book not found")
	}
	bs.db.Model(&bookToUpdate).Updates(book)

	return nil
}

