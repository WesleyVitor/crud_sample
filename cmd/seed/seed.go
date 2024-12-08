package main

import (
	"bufio"
	"example/work-at-olist-challenge/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
func main(){
	fmt.Println("Initializing seed")
	db, err := gorm.Open(sqlite.Open("olist_db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&models.Author{})

	file, err := os.Open("cmd/seed/authors.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
		if scanner.Text() != "name" && db.Find(&models.Author{}, "name = ?", scanner.Text()).RowsAffected == 0 {
				author := models.Author{Name: scanner.Text()}
				db.Create(&author)
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}