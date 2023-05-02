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

//go:embed views/*.gohtml
var goHtml embed.FS
var goTmpl = template.Must(template.ParseFS(goHtml, "views/*.gohtml"))

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	if err := goTmpl.ExecuteTemplate(w, "autoEscape.gohtml", map[string]interface{}{
		"Title": "Auto Escape",
		"Body":  template.HTMLEscapeString("<h1>Hallo Guys, welcome back to my channel</h1>"), // will still be displayed like this

	}); err != nil {
		panic(err)
	}
}

func TestAutoEscapeI(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestAutoEscapeServer(t *testing.T) {
	mux := new(http.ServeMux)
	mux.HandleFunc("/", TemplateAutoEscape)

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux
	server.ListenAndServe()
}
