package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get greeting from environment variable
	greeting := os.Getenv("GREETING")
	// Write GREETING, World to the http response
	fmt.Fprintf(w, "%s, World", greeting)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
