package golangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ======================== //
// template caching = menyimpan template ke dalam global variable
// ======================== //

//go:embed templates/*.gohtml
var templatesCache embed.FS

var myTemplate = template.Must(template.ParseFS(templatesCache, "templates/*.gohtml"))

func TemplateCache(w http.ResponseWriter, r *http.Request) {
	t := myTemplate
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Directory Template")
}

func TestTemplateCache(t *testing.T) {
	request :=httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCache(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}