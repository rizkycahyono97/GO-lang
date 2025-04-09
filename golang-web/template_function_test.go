package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ======================== //
// golang template function = membuat function didalam template
// ======================== //
type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string{
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello  "Rizky Cahyono"}}`))	// function SayHello dengan parameter
	
	// struct
	data := MyPage{
		Name: "Asep Septiawan",
	}

	t.ExecuteTemplate(w, "FUNCTION", data)
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}



// ======================== //
// global function = function bawaan golang
// ======================== //
func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len  "Rizky Cahyono"}}`))	// global function len = untuk menghitung panjang string
	
	// struct
	data := MyPage{
		Name: "Asep Septiawan",
	}

	t.ExecuteTemplate(w, "FUNCTION", data)
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}



// ====================== //
// membuat global function sendiri
// ====================== //
func TemplateGlobalfunction(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")

	// registrasi global function yg kita buat
	t = t.Funcs(map [string]interface{} {
		"toUpper": func (value string) string {	// function toUppercase
			return strings.ToUpper(value)
		},
	})

	data := MyPage{
		Name: "Rizky Cahyono",
	}

	t = template.Must(t.Parse(`{{ toUpper .Name }}`))

	t.ExecuteTemplate(w, "FUNCTION", data)
}

func TestTemplateGlobalFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateGlobalfunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}


// ========================= //
// function global pipeline = untuk menghubungkan function dg pipeline |
// ========================= //
func TemplateGlobalfunctionPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")

	// registrasi global function yg kita buat
	t = t.Funcs(map [string]interface{} {
		"sayHello":  func (name string) string {
			return "Hello " + name
		},
		"toUpper": func (value string) string {	// function toUppercase
			return strings.ToUpper(value)
		},
	})

	data := MyPage{
		Name: "rizky cahyono putra",
	}

	t = template.Must(t.Parse(`{{ sayHello .Name | toUpper }}`)) 	// function sayHello dihubungkan dg toUpper

	t.ExecuteTemplate(w, "FUNCTION", data)
}

func TestTemplateGlobalFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateGlobalfunctionPipeline(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}