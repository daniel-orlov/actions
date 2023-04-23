package transport

import (
	"log"
	"net/http"
)

// Handler returns an http.HandlerFunc that can be used to handle a request.
// It is a simple example of an HTTP handler that writes a response.
func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello Actions"))
		if err != nil {
			log.Printf("Error writing response: %v", err)
			return
		}
	}
}
