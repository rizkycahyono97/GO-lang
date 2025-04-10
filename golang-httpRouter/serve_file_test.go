package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// ========================= //
// Serve File
// ========================= //
// Serve File = untuk merender file statis seperti html, css, js, dll


//go:embed resources
var resources embed.FS
func TestServeFile(t *testing.T) {

	router := httprouter.New()	 
	directory, _ := fs.Sub(resources, "resources")	// untuk menghapus http://localhost:8080/resources menjadi http://localhost:8080/files
	router.ServeFiles("/files/*filepath", http.FS(directory))

	// request
	request := httptest.NewRequest("GET", "http://localhost:8080/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	// Serve HTTP from httprouter
	router.ServeHTTP(recorder, request)

	// Check the response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello HTTP Router", string(body))
}

func TestServeFileGoodBye(t *testing.T) {

	router := httprouter.New()	 
	directory, _ := fs.Sub(resources, "resources")	// untuk menghapus http://localhost:8080/resources menjadi http://localhost:8080/files
	router.ServeFiles("/files/*filepath", http.FS(directory))

	// request
	request := httptest.NewRequest("GET", "http://localhost:8080/files/goodbye.txt", nil)
	recorder := httptest.NewRecorder()

	// Serve HTTP from httprouter
	router.ServeHTTP(recorder, request)

	// Check the response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Good Bye HTTP Router", string(body))
}