package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Print("Hello, World!")

	router := mux.NewRouter()
	router.HandleFunc("/", Hello)

	err := http.ListenAndServe(":8888", router)

	if err != nil {
		log.Fatal(err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
