package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	s := http.Server{Addr: ":8080", Handler: mux}
	mux.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello")
	})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		s.Shutdown(context.Background())
	})
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Printf("Finished")
}
