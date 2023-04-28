package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	if name := r.PostForm.Get("name"); name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Please provide name")
	}

	if address := r.PostForm.Get("address"); address == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Please provide address")
	}

}

func TestResponseCodeInvalid(t *testing.T) {
	dataBody := strings.NewReader("name=Oldest&address=")
	request := httptest.NewRequest("POST", "http://localhost:8080/post", dataBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	ResponseCode(recorder, request)
	result := recorder.Result()

	data, _ := io.ReadAll(result.Body)
	stringData := string(data)
	fmt.Println(stringData)
	fmt.Println("Response code:", result.StatusCode)
	fmt.Println(result.Status)

}

func TestResponseCodeValid(t *testing.T) {
	dataBody := strings.NewReader("name=Oldest&address=Mejad")
	request := httptest.NewRequest("POST", "http://localhost:8080/post", dataBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	ResponseCode(recorder, request)
	result := recorder.Result()

	data, _ := io.ReadAll(result.Body)
	stringData := string(data)
	fmt.Println(stringData)
	fmt.Println("Response code:", result.StatusCode)
	fmt.Println(result.Status)

}
