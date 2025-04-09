package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ========================= //
// cookie
// ========================= //

// set Cookie
func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success set Cookie")
}

// get Cookie
func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(w, "Cookie not Found")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

// server test
func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// how to run
// http://localhost:8080/set-cookie?name=rizky-cahyono-putra
// http://localhost:8080/get-cookie


// test setCookie
func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/set-cookie?name=Rizky-Cahyono-Putra", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("cookie %s %s \n", cookie.Name, cookie.Value)
	}
}

// test getCookie
func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)

	cookie.Name = "X-PZN-Name"
	cookie.Value = "Rizky-Cahyono-Putra"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}