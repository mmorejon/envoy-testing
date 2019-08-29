package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Create a server on port 8000
	// Exactly how you would run an HTTP/1.1 server
	srv := &http.Server{Addr: ":8000", Handler: http.HandlerFunc(handle)}

	// Start the server with TLS, since we are running HTTP/2 it must be
	// run with TLS.
	// Exactly how you would run an HTTP/1.1 server with TLS connection.
	log.Printf("Serving on https://0.0.0.0:8000")
	log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
	// log.Fatal(srv.ListenAndServe())
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Get hostname
	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v", err)
	}
	// Log the request protocol
	log.Printf("Got connection: %s - %s", r.Proto, name)
	// Send a message back to the client
	w.Write([]byte("Hello - " + name))
}
