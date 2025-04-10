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
// Method Not Allowed Handler 
// ========================= //
// Test Jika Method HTTP yang dikirim user tidak cocok dengan yang ada di router
func TestMethodNotAllowed(t *testing.T) {

	router := httprouter.New()
	
	// Not Found Handler
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Boleh Harus Pakai POST")
	})

	// POST Handler
	router.POST("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
		fmt.Fprint(w, "OK")
	})

	// request dengan GET	
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	// Serve HTTP from httprouter
	router.ServeHTTP(recorder, request)

	// Check the response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Boleh Harus Pakai POST", string(body))
}