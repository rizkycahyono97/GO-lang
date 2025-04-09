package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ==================== //
// query parameter
// ==================== //

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParamater(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/hello?name=rizky", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}


// ==================== //
// multiple parameter
// ==================== //
func MultipleParameter(w http.ResponseWriter, r *http.Request) {
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")

	if first_name  == "" {
		fmt.Fprint(w, "Hello")
	} else if last_name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s %s", first_name, last_name)
	}
}

func TestMultipleParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/hello?first_name=rizky&last_name=cahyono", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}



// ======================= //
// multiple parameter values
// ======================= //
func MultipleParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]	// parameter
	fmt.Fprint(w, strings.Join(names, " "))		// menggabungkan string dg spasu
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/hello?name=rizky&name=cahyono&name=putra", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

