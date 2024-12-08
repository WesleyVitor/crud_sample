package services

import (
	"errors"
	"example/work-at-olist-challenge/internal/models"

	"gorm.io/gorm"
)

type AuthorService struct {
	db *gorm.DB
}

func NewAuthorService(db *gorm.DB) AuthorService {
	return AuthorService{
		db: db,
	}
}

func (as AuthorService) GetAuthorsByIDs(ids []int) ([]*models.Author, error) {
	var authors []*models.Author
	as.db.Find(&authors, ids)
	if len(authors) != len(ids) {
		return nil, errors.New("some authors were not found")
	}
	return authors, nil
}