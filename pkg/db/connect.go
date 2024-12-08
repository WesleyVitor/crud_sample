package db

import (
	"example/work-at-olist-challenge/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// Code for connecting to the database
	db, err := gorm.Open(sqlite.Open("olist_db.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Author{})

	return db
}
