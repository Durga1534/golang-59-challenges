package models

import (
	"fmt"

	"github.com/Durga1534/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Init() {
	config.Connect()
	db = config.GetDB()
	if db != nil {
		db.AutoMigrate(&Book{})
		fmt.Println("Database initialized successfully")
	}
}

func (b *Book) CreateBook() *Book {
	if db == nil {
		fmt.Println("Error: Database not connected")
		return b
	}
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	if db == nil {
		fmt.Println("Error: Database not connected")
		return books
	}
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	if db == nil {
		fmt.Println("Error: Database not connected")
		return &getBook, nil
	}
	result := db.Where("ID=?", id).Find(&getBook)
	return &getBook, result
}

func (b *Book) UpdateBook() *Book {
	if db == nil {
		fmt.Println("Error: Database not connected")
		return b
	}
	db.Save(&b)
	return b
}

func DeleteBook(id int64) Book {
	var book Book
	if db == nil {
		fmt.Println("Error: Database not connected")
		return book
	}
	db.Where("ID=?", id).Delete(&book)
	return book
}
