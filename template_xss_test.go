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

func TestArraySign(t *testing.T) {
	nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}

	multiply := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			fmt.Println(0)
		}

		if nums[i] < 0 {
			multiply *= -1
		}
	}
	fmt.Println(multiply)

}
