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
// Router Pattern By Name (:)
// ========================= //
// pattern name = params yang menggunakan nama sebagai parameter
// pattern ini menggunakan tanda : pada akhir url

func TestRouterParams(t *testing.T) {

	router := httprouter.New()

	// menggunakan parameter id
	router.GET("/products/:id/items/:itemId", func (writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + id + " Item " + itemId 
		fmt.Fprint(writer, text)
	})

	// request
	request := httptest.NewRequest("GET", "http://localhost:8080/products/1/items/2", nil)
	recorder := httptest.NewRecorder()

	// Serve HTTP from httprouter
	router.ServeHTTP(recorder, request)

	// Check the response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// hasil
	assert.Equal(t, "Product 1 Item 2", string(body))
}



// ========================= //
// Router Pattern By File (*)
// ========================= //
// pattern file = params yang menggunakan nama file sebagai parameter
// pattern ini menggunakan tanda * pada akhir url

func TestRouterPatternCatchAllParameter(t *testing.T) {

	router := httprouter.New()

	// menggunakan parameter id
	router.GET("/images/*image", func (writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
		image := params.ByName("image")
		text := "Image : " + image
		fmt.Fprint(writer, text)
	})

	// request
	request := httptest.NewRequest("GET", "http://localhost:8080/images/small/profile.jpeg", nil)
	recorder := httptest.NewRecorder()

	// Serve HTTP from httprouter
	router.ServeHTTP(recorder, request)

	// Check the response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// hasil
	assert.Equal(t, "Image : /small/profile.jpeg", string(body))
} 