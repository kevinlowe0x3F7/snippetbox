package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func main() {
	// register route
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// starting server
	log.Print("starting server on: 4000")

	// ListenAndServe to start the server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
