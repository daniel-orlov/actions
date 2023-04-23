package main

import (
	"log"
	"net/http"

	"github.com/daniel-orlov/actions/transport"
)

func main() {
	// Create a simple server that listens on port 8080 and
	// serves "Hello Actions" on the root path.
	log.Fatal(http.ListenAndServe(":8080", transport.Handler()))
}
