package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ======================== //
// template layout = seperti blade template engine di laravel
// ======================== //
func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./template_layout/header.gohtml",
		"./template_layout/layout.gohtml",
		"./template_layout/footer.gohtml",
		))

	data := map[string]interface{} {
		"Title": "Template Layout",
		"Name": "Rizky Cahyono Putra",
	}
	
	t.ExecuteTemplate(w, "layout", data)
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
