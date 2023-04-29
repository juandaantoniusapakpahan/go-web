package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Student struct {
	Name  string
	Grade int
}

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./views/if.html")
	if err != nil {
		panic(err)
	}

	data := Student{
		Name:  "Gorgioa",
		Grade: 75,
	}

	if err := tmpl.ExecuteTemplate(w, "if.html", data); err != nil {
		panic(err)
	}
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/action-if", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./views/for.html")
	if err != nil {
		panic(err)
	}

	data := []Student{
		{
			Name:  "Jhonson",
			Grade: 60,
		},
		{
			Name:  "Rin",
			Grade: 90,
		},
	}
	if err := tmpl.ExecuteTemplate(w, "for.html", data); err != nil {
		panic(err)
	}
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/action-for", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
