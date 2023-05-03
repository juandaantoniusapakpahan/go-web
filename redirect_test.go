package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "redirect-to", http.StatusTemporaryRedirect)
}

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello to Redirect to")
}

func TestRedirectServers(t *testing.T) {
	mux := new(http.ServeMux)
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
