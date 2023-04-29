package goweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	if params := r.URL.Query().Get("name"); params != "" {
		http.ServeFile(w, r, "./assets/index.html")
	} else {
		http.ServeFile(w, r, "./assets/notFound.html")
	}
}

func TestServeFile(t *testing.T) {
	mux := new(http.ServeMux)
	mux.HandleFunc("/", ServeFileHandler)
	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux
	server.ListenAndServe()
}

//go:embed assets/index.html

var index string

//go:embed assets/notFound.html

var notFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if param := r.URL.Query().Get("name"); param != "" {
		fmt.Fprint(w, index)
	} else {
		fmt.Fprint(w, notFound)
	}
}

func TestServerFileEmbed(t *testing.T) {
	mux := new(http.ServeMux)
	mux.HandleFunc("/", ServeFileEmbed)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
