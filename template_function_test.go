package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TmplStd struct {
	Name string
}

func (tmpl TmplStd) ReturnName(name string) string {
	return "Hello " + name + ". " + "I am " + tmpl.Name + ". Nice to meet You"
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("TFuction").Parse(`{{.ReturnName "Rojoko"}}`))

	data := TmplStd{Name: "Hologram"}
	if err := tmpl.ExecuteTemplate(w, "TFuction", data); err != nil {
		panic(err)
	}
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/template-funciton", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func TemplateFunctionGlob(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("FUNCTION")
	tmpl = tmpl.Funcs(map[string]interface{}{
		"upper": func(name string) string {
			return strings.ToUpper(name)
		},
		"lower": func(name string) string {
			return strings.ToLower(name)
		},
	})
	tmpl = template.Must(tmpl.Parse(`{{upper .Name}}, {{lower .Name}}`))
	if err := tmpl.ExecuteTemplate(w, "FUNCTION", TmplStd{Name: "Zoro Dono"}); err != nil {
		panic(err)
	}
}

func TestFunctionGlob(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/func-glob", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlob(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func TemplateFunctionGlobPipeline(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("FUNCTION")
	tmpl = tmpl.Funcs(map[string]interface{}{
		"upper": func(name string) string {
			return strings.ToUpper(name)
		},
		"lower": func(name string) string {
			return strings.ToLower(name)
		},
		"helloSay": func(name string) string {
			return "Hello " + name
		},
	})
	tmpl = template.Must(tmpl.Parse(`{{helloSay .Name | upper }}`))
	if err := tmpl.ExecuteTemplate(w, "FUNCTION", TmplStd{Name: "Zoro Dono"}); err != nil {
		panic(err)
	}
}

func TestFunctionGlobPipeline(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/func-glob", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobPipeline(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
