package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mlukanova/go-workshop/internal/diagnostics"
)

func main() {
	log.Print("Starting Application")

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
		log.Print("The Application Server is preparing")

		err := http.ListenAndServe(":"+blPort, router)

		if err != nil {
			log.Fatal(err)
		}
	}()

	diagnostics := diagnostics.NewDiagnostics()
	log.Print("The Diagnostics Server is preparing")

	err := http.ListenAndServe(":"+diagnosticsPort, diagnostics)

	if err != nil {
		log.Fatal(err)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Print("The hello function was called")
	w.WriteHeader(200)
}
