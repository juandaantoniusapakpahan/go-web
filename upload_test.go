package goweb

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

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	name := r.PostFormValue("name")
	if err != nil {
		panic(err)
	}

	fileDistination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDistination, file)
	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "upload.success.html", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}

func IdxHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "upload.form.html", nil); err != nil {
		panic(err)
	}
}

func TestUploadServer(t *testing.T) {
	mux := new(http.ServeMux)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))
	mux.HandleFunc("/", IdxHandler)
	mux.HandleFunc("/upload", UploadFile)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

//go:embed resources/NY.jpg
var uploadFileIMG []byte

func TestFileUpload(t *testing.T) {
	body := new(bytes.Buffer)

	write := multipart.NewWriter(body)
	write.WriteField("name", "Richardo Albalasta")
	file, _ := write.CreateFormFile("file", "picture.jpg")
	file.Write(uploadFileIMG)
	write.Close()

	request := httptest.NewRequest("POST", "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", write.FormDataContentType())
	recorder := httptest.NewRecorder()

	UploadFile(recorder, request)

	responseBody, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(responseBody))
}
