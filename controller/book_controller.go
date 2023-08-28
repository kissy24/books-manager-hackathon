package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kissy24/go-reading-books/model"
)

func ShowAllBooks(ctx *gin.Context) {
	book := model.SelectAll()
	ctx.JSON(200, gin.H{"book": book})
}

func AddBook(ctx *gin.Context) {
	var book model.Book
	ctx.BindJSON(&book)
}
