package services

import (
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


