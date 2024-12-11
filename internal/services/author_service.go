package services

import (
	"errors"
	"example/work-at-olist-challenge/internal/models"

	"gorm.io/gorm"
)

type AuthorService struct {
	db *gorm.DB
}

// Create a new AuthorService instance
func NewAuthorService(db *gorm.DB) AuthorService {
	return AuthorService{
		db: db,
	}
}

// Get authors by IDs and return an error if some author is not found
func (as AuthorService) GetAuthorsByIDs(ids []int) ([]*models.Author, error) {
	var authors []*models.Author
	as.db.Find(&authors, ids)
	if len(authors) != len(ids) {
		return nil, errors.New("some authors were not found")
	}
	return authors, nil
}