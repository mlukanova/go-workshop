package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlukanova/go-workshop/internal/diagnostics"
)

func main() {
	log.Print("Hello, World!")

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	go func() {
		err := http.ListenAndServe(":8888", router)

		if err != nil {
			log.Fatal(err)
		}
	}()

	diagnostics := diagnostics.NewDiagnostics()
	err := http.ListenAndServe(":8885", diagnostics)

	if err != nil {
		log.Fatal(err)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
