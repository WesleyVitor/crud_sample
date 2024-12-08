package services

import (
	"example/work-at-olist-challenge/pkg/models"

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

func (as AuthorService) GetAuthorsByIDs(ids []int) []*models.Author {
	var authors []*models.Author
	as.db.Find(&authors, ids)
	return authors
}