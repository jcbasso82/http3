package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	// Set up HTTP server multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)

	// Log the server start
	fmt.Println("Starting HTTP/3 server on https://localhost:4242")

	// Start listening and serving using HTTP/3
	err := http3.ListenAndServeTLS(":4242", "./cert.pem", "./key.pem", mux)

	if err != nil {
		log.Fatalf("Failed to configure HTTP/3: %v", err)
	}

}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pong"}`))
}
