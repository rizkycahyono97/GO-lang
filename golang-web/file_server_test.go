package golangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// ======================== //
// file server
// ======================== //
// untuk merender static file seperti html, css, js, image

func TestFileServer(t *testing.T) {
	
	// directory
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	// http handler
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// server
	server := http.Server{
		Addr:	"localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// how to access
// http://localhost:8080/static/index.html
// http://localhost:8080/static/css/style.css
// http://localhost:8080/static/js/script.js



// ========================= //
// file server with golang embed
// ========================= //

//go:embed resources/*
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	// golang embed
	directory, _ := fs.Sub(resources, "resources")	// supaya bisa langsung akses ke static/ tanpa static/resource
	fileServer := http.FileServer(http.FS(directory))

	// http handler
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// server
	server := http.Server{
		Addr:	"localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// how to access
// http://localhost:8080/static/index.html
// http://localhost:8080/static/css/style.css
// http://localhost:8080/static/js/script.js