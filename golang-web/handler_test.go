package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// =================== //
// handler = response
// =================== //
func TestHandler(t *testing.T) {
	
	// response
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {	// w = response
		fmt.Fprint(w, "Hello World")
	}

	// server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


// ===================== //
// serveMux = seperti Router
// ===================== //
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	// route / endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Root Directory")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Hi Directory")
	})

	// url pattern
	// karena kita memakai slash dibelakang images/ , maka jika kita menulis images/asdf maka 
	// yang di render akan tetap di /images/, bukan malah not found
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Images")
	})

	mux.HandleFunc("/images/thumbnail/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Thumbnail")
	})

	// server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}



// ================= //
// request
// ================= //
func TestRequst(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)		// mengetahui method http yang dipakai
		fmt.Fprintln(w, r.RequestURI)	// menampilkan request URL
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}
	
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

