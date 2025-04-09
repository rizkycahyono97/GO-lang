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
// template auto escape = untuk menghindari XSS
// ======================== //
func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body": "<script>alert('Hello')</script>",	// akan otomatis di escape / tidak dirender sebagai tag HTML	
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// test server
func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


// ======================== //
// mematikan auto escape
// ======================== //
func TemplateAutoNoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto NoEscape",
		"Body": template.HTML("<script>alert('Hello No Escape')</script>"),	// akan otomatis di No Escape / tag HTML,CSS,JS akan dirender
	})
}

func TestTemplateAutoNoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoNoEscape(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// test server
func TestTemplateAutoNoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoNoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


// ======================== //
// contoh XSS di Parameter
// ======================== //
func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Percobaan XSS",
		"Body": template.HTML(r.URL.Query().Get("body")),	// akan otomatis di No Escape / tag HTML,CSS,JS akan dirender
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?body=<script>alert('XSS')</script>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// test server
func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}