package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Open DB
func Open() *gorm.DB {
	db, err := gorm.Open("sqlite3", "xml.sqlite3")
	if err != nil {
		panic("Don't open sqlite3")
	}
	return db
}
