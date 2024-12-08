package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string
	Books []*Book `gorm:"many2many:author_book;"`
} 

type Book struct {
	gorm.Model
	Name string 
	Edition int
	PublicationYear int
	Authors []*Author `gorm:"many2many:author_book;"`

}
