package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kissy24/go-reading-books/model"
)

func ShowAllBooksAPI(ctx *gin.Context) {
	book := model.SelectAll()
	ctx.JSON(200, gin.H{"book": book})
}

func ShowAllBooks(ctx *gin.Context) {
	book := model.SelectAll()
	ctx.HTML(200, "index.html", gin.H{"book": book})
}

func AddBookAPI(ctx *gin.Context) {
	var book model.Book
	ctx.BindJSON(&book)
}

func AddBook(ctx *gin.Context) {
	var form model.Book
	// ここがバリデーション部分
	if err := ctx.Bind(&form); err != nil {
		book := model.SelectAll()
		ctx.HTML(http.StatusBadRequest, "index.html", gin.H{"book": book, "err": err})
		ctx.Abort()
	} else {
		title := ctx.PostForm("title")
		isbn, err := strconv.Atoi(ctx.PostForm("isbn"))
		if err != nil {
			book := model.SelectAll()
			ctx.HTML(http.StatusBadRequest, "index.html", gin.H{"book": book, "err": err})
			ctx.Abort()
		}
		publisher := ctx.PostForm("publisher")
		book := model.Book{Title: title, ISBN: isbn, Publisher: publisher}
		model.Insert(book)
		ctx.Redirect(302, "/")
	}
}
