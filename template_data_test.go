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

var viewData embed.FS

type M map[string]interface{}
type Website struct {
	Title, Name string
}

func TemplateData(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(viewData, "views/*.html")
	if err != nil {
		panic(err)
	}
	data := M{
		"Title": "Template Data",
		"Name":  "Jhon Rock",
	}
	if err := tmpl.ExecuteTemplate(w, "templateData.html", data); err != nil {
		panic(err)
	}
}

func TestTemplateData(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/template-data", nil)
	recorder := httptest.NewRecorder()

	TemplateData(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(viewData, "views/*.html")
	if err != nil {
		panic(err)
	}
	data := Website{
		Title: "JJJMAIL",
		Name:  "Rinnata",
	}
	if err := tmpl.ExecuteTemplate(w, "templateData.html", data); err != nil {
		panic(err)
	}
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/template-data", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
