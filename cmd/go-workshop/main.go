package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Hello, World!")

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		log.Fatal(err)
	}
}
