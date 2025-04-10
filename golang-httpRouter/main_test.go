package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
)

// ========================= //
// a simple http router
// ========================= //
func TestMain(t *testing.T) {
	// Initialize the package httprouter
	router := httprouter.New()
	// simple handler with httprouter GET
	router.GET("/", func (writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
		fmt.Fprint(writer, "Hello HTTP ROUTER")
	})

	server := http.Server{
		Handler: router,
		Addr: "localhost:8080",
	}

	server.ListenAndServe()
}