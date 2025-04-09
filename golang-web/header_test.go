package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ======================== //
// header
// ======================== //
func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")	// get header 
	fmt.Fprint(w, contentType)
}
func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	
	// set header
	request.Header.Set("content-type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body,  _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}


// ========================= //
// response header
// ========================= //
func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Rizky Cahyono Putra")
	fmt.Fprint(w, "Ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	
	// set header
	request.Header.Set("content-type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body,  _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-by"))
}