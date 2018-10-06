package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/mlukanova/go-workshop/internal/diagnostics"
)

type serverConf struct {
	port   string
	name   string
	router http.Handler
}

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

	diagnostics := diagnostics.NewDiagnostics()
	possibleErrors := make(chan error, 2)

	configurations := []serverConf{
		{
			port:   blPort,
			router: router,
			name:   "application server",
		},
		{

			port:   diagnosticsPort,
			router: diagnostics,
			name:   "diagnostics server",
		},
	}

	servers := make([]*http.Server, 2)

	for i, c := range configurations {
		go func(conf serverConf, i int) {
			log.Printf("The %s is preparing to handle connections...", conf.name)
			servers[i] = &http.Server{
				Addr:    ":" + conf.port,
				Handler: conf.router,
			}
			err := servers[i].ListenAndServe()
			if err != nil {
				possibleErrors <- err
			}
		}(c, i)
	}

	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-possibleErrors:
		log.Printf("Got an error: %v", err)
	case sig := <-interupt:
		log.Printf("Got signal: %v", sig)
	}

	for _, s := range servers {
		timeout := 5 * time.Second
		log.Printf("Shutdown with timeout: %s", timeout)
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err := s.Shutdown(ctx)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf("Server gracefully stopped")
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Print("The hello function was called")
	w.WriteHeader(200)
}
