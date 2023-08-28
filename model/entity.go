package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title     string
	ISBN      int
	Publisher string
}
