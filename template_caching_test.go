package goweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed views/*.html

var templCaching embed.FS

var tmpl = template.Must(template.ParseFS(templCaching, "views/*.html"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "caching.html", nil); err != nil {
		panic(err)
	}
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/template-parse", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
