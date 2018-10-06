package diagnostics

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewDiagnostics() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", health)
	router.HandleFunc("/info", ready)
	return router
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
