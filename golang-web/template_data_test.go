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
// golang template data
// ======================== //
func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{		// mengggunakan tipe data map
		"Title": "Template Data Map",
		"Name": "Rizky Cahyono Putra",
		"Address": map[string]interface{}{	// nested map
			"Street": "Jl. Raya",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// map
type Address struct {
	Street string
	City string
}

type Page struct {
	Title string
	Name string
	Address Address
}

// menggunakan struct
func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name: "Rizky Cahyono Putra",
		Address: Address {	// nested struct
			Street: "Jl. Raya",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

