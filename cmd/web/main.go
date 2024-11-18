package main

import (
	"log"
	"net/http"
)

func main() {
	// register route
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// starting server
	log.Print("starting server on: 4000")

	// ListenAndServe to start the server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
