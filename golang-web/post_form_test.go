package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ================== //
// post form
// ================== //

// handler
func PostFormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("first_name=Rizky&last_name=Cahyono")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recoreder := httptest.NewRecorder()

	PostFormHandler(recoreder, request)

	response := recoreder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}