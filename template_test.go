package goweb

import (
	"embed"
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

func TemplateParseFile(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/about.html"))

	if err := tmpl.ExecuteTemplate(w, "about.html", "Hello About"); err != nil {
		panic(err)
	}
}

func TestTemplateParseFile(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/about", nil)
	recorder := httptest.NewRecorder()

	TemplateParseFile(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateGlobalParseFile(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseGlob("./views/*.html")
	if err != nil {
		panic(err)
	}

	if err := tmp.ExecuteTemplate(w, "about.html", "Hello About"); err != nil {
		panic(err)
	}
}

func TestTemplateGlobalParseFile(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/about", nil)
	recorder := httptest.NewRecorder()

	TemplateGlobalParseFile(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

//go:embed views/*.html

var views embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(views, "views/*.html")
	if err != nil {
		panic(err)
	}
	if err := tmpl.ExecuteTemplate(w, "dashboard.html", "Welcome to Dashboard"); err != nil {
		panic(err)
	}
}

func TestTempateEmbed(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/about", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
