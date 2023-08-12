package main

import (
	"fmt"
	"net/http"
)

func main() {
	h := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello")
	}

	http.HandleFunc("/", h)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
