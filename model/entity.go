package model

type Book struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	ISBN      uint   `json:"isbn"`
	Publisher string `json:"publisher"`
}
