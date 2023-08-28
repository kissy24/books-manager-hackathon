package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("Don't open sqlite3")
	}
	db.LogMode(true)
	db.AutoMigrate(&Book{})
	return db
}

func SelectAll() (books []Book) {
	db := Init()
	result := db.Find(&books)
	if result.Error != nil {
		panic(result.Error)
	}
	defer db.Close()
	return
}

func Insert(book Book) {
	db := Init()
	db.Create(&book)
	defer db.Close()
}
