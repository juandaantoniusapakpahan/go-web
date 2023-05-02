package goweb

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed views/*.html

var lyt embed.FS

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(lyt, "views/*.html"))
	if err := tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Name":   "Richard",
		"Title":  "Richard Dev",
		"Footer": "@Richard",
	}); err != nil {
		panic(err)
	}
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/tmpl-layout", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
