package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)

	cookie.Name = "cookie"
	cookie.Value = "Holaldro Smith"
	cookie.Path = "/"
	cookie.HttpOnly = true

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Hellow ", cookie.Value)
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("cookie")
	if err != nil {
		fmt.Fprint(w, "No cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

func TestCookieServer(t *testing.T) {
	mux := new(http.ServeMux)
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux
	server.ListenAndServe()
}

func TestCookies(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/set-cookie", nil)
	recoreder := httptest.NewRecorder()

	SetCookie(recoreder, request)
	result := recoreder.Result()
	data, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	stringData := string(data)
	cookies := result.Cookies()
	for _, cookie := range cookies {
		assert.NotEqual(t, cookie.Name, nil)
		assert.NotEqual(t, cookie.Value, nil)
	}
	fmt.Println(stringData)
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/get-cookie", nil)
	cookie := new(http.Cookie)
	cookie.Name = "cookie"
	cookie.Value = "Roro zororo"
	cookie.Path = "/"
	cookie.HttpOnly = true
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	cookies := response.Cookies()
	data, _ := io.ReadAll(response.Body)
	fmt.Println(string(data))
	for _, cookie := range cookies {
		assert.NotEqual(t, cookie.Name, nil)
		assert.NotEqual(t, cookie.Value, nil)
	}
}
