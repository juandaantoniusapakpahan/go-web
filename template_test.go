package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	htmlText := `<html><body>{{.}}</body></html>`
	html := template.Must(template.New("index").Parse(htmlText))

	if err := html.ExecuteTemplate(w, "index", "Hello, I am Marialina"); err != nil {
		panic(err)
	}
}

func TestTemplateHtml(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/html", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
