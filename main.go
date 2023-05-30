package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nb, err := fmt.Fprintf(w, "hellow browser")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("bytes: ", nb)

	})
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
