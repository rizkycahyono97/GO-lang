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
// Panic Handler
// ========================= //
func TestPanic(t *testing.T) {

	router := httprouter.New()

	// panic handler
	router.PanicHandler = func (w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(w, "Error : ", error)
	}

	router.GET("/", func (writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
		panic("Ups")
	})

	// request
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	// Serve HTTP from httprouter
	router.ServeHTTP(recorder, request)

	// Check the response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Error : Ups", string(body))
}