package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	name := r.PostForm.Get("name")
	address := r.PostForm.Get("address")

	fmt.Fprint(w, name, " ", address)
}

func TestFormPost(t *testing.T) {
	bodyForm := strings.NewReader("name=Golang&address=USA")
	request := httptest.NewRequest("POST", "http://localhost:8080/post", bodyForm)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recoreder := httptest.NewRecorder()

	FormPost(recoreder, request)
	result := recoreder.Result()
	data, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	stringData := string(data)
	fmt.Println(stringData)
}
