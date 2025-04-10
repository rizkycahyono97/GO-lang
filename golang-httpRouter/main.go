package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// Initialize the package httprouter
	router := httprouter.New()

	server := http.Server{
		Handler: router,
		Addr: "localhost:8080",
	}

	server.ListenAndServe()
}
