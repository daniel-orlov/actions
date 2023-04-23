package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a simple server that listens on port 8080 and
	// serves "Hello Actions" on the root path.
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello Actions"))
		if err != nil {
			log.Fatal(err)
		}
	})))
}
