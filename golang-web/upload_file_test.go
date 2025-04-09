package golangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// render upload.form.gohtml
func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "upload.form.gohtml", nil)
}

// handle upload
func Upload(w http.ResponseWriter, r *http.Request) {

	// parse form dahulu
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	// destination file
	fileDestination, err := os.Create("./storage/" + fileHeader.Filename)	// fileHeader.fileName = nama file
	if err != nil {
		panic(err)
	}

	// copy file dari file ke fileDestination
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	// sekarang tinggal yang name
	name := r.PostFormValue("name")
	myTemplate.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,	// static = file disimpan di folder static sementara
	})
}

// server upload
func TestServerUploadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./storage"))))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


// ======================== //
// test upload file = implementasi unit testnya
// ======================== //
//go:embed storage/1.jpeg
var uploadFileTest []byte
func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Rizky Cahyono Putra")
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.jpeg")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}