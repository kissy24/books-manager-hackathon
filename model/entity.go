package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	ID        int    `json:"id"`
	Title     string `json:"title"`
	ISBN      int    `json:"isbn"`
	Publisher string `json:"publisher"`
}
