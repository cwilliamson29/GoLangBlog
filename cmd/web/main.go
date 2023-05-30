package main

import (
	"github.com/cwilliamson29/GoLangBlog/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
