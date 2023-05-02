package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	body := r.URL.Query().Get("body")
	if err := goTmpl.ExecuteTemplate(w, "autoEscape.gohtml", map[string]interface{}{
		"Body":  body,
		"Title": "XSS",
	}); err != nil {
		panic(err)
	}
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateXSS(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestSeverSXX(t *testing.T) {
	mux := new(http.ServeMux)
	mux.HandleFunc("/", TemplateXSS)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
