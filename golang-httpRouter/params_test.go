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
// params
// ========================= //
func TestParams(t *testing.T) {

	router := httprouter.New()

	// menggunakan parameter id
	router.GET("/products/:id", func (writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
		id := params.ByName("id")
		text := "Product " + id 
		fmt.Fprint(writer, text)
	})

	// request
	request := httptest.NewRequest("GET", "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	// Serve HTTP from httprouter
	router.ServeHTTP(recorder, request)

	// Check the response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// hasil
	assert.Equal(t, "Product 1", string(body))
}