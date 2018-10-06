package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mlukanova/go-workshop/internal/diagnostics"
)

func main() {
	log.Print("Hello, World!")

	blPort := os.Getenv("GO_PORT")

	if len(blPort) == 0 {
		log.Fatal("Config port should be set!")
	}

	diagnosticsPort := os.Getenv("GO_DIAGNOSTICS_PORT")
	if len(diagnosticsPort) == 0 {
		log.Fatal("Diagnostics port should be set!")
	}

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
