package goweb

import (
	"fmt"
	"net/http"
	"testing"
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
