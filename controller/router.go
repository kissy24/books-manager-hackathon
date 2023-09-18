package controller

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*.html")
	r.GET("/", ShowAllBooks)
	r.POST("/add", AddBook)
	r.GET("/api", ShowAllBooksAPI)
	r.POST("/add/api", AddBookAPI)
	return r
}
