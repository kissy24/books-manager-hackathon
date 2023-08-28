package main

import "github.com/kissy24/go-reading-books/controller"

func main() {
	router := controller.GetRouter()
	router.Run()
}
