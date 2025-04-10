package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)


// ========================= //
// a simple http router
// ========================= //
func TestRouter(t *testing.T) {

	router := httprouter.New()
	router.GET("/", func (writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
		fmt.Fprint(writer, "Hello World")
	})

	// request
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	// Serve HTTP from httprouter
	router.ServeHTTP(recorder, request)

	// Check the response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(body))
}