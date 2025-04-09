package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// ==================== //
// download file
// ==================== //
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")

	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "File Not Found")
		return	// break
	} 

	// supaya ketika di download, tidak perlu/udah dirender
	w.Header().Set("Content-Disposition", "attachment; filename=\""+file+"\"")

	// file disimpan di folder storage
	http.ServeFile(w, r, "./storage/" + file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// cara run
// localhost:8080?file=1.jpeg