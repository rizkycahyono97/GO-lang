package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// ================== //
// middleware = log
// ================== //
//Middleware di Go (Golang) adalah fungsi/perantara yang berjalan sebelum atau sesudah sebuah Handler dieksekusi.
type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(w, r) // panggil handler yang ada di middleware
	fmt.Println("After Execute Handler")
}

//  ================== //
// contoh middleware di error handler
//  ================== //
type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)

}

// Test
func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	// contoh handler utama
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Handler Executed")
		fmt.Fprint(w, "Hello Middleware")
	})

	// contoh handler foo
	mux.HandleFunc("/foo", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Foo Executed")
		fmt.Fprint(w, "Hello Foo")
	})

	// contoh handler panic/error handler
	mux.HandleFunc("/panic", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Foo Executed")
		panic("Ups")
	})

	// log middleware
	LogMiddleware := &LogMiddleware{
		Handler: mux,
	}

	// error handler middleware
	ErrorHandler := &ErrorHandler{
		Handler: LogMiddleware,
	}

	// server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: ErrorHandler,		// di handler ErrorHandler->LogMiddleware->mux
	}
	
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// alur eksekusi
// 1. request server
// 2. di handle oleh LogMiddleware
// 3. di handle oleh mux
// 4. di handle oleh handler / handler utama
// 5. di handle lagi oleh LogMiddleware



