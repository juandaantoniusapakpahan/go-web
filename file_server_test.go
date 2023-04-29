package goweb

import (
	"embed"
	_ "embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	mux := new(http.ServeMux)

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./assets"))))
	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed assets

var assets embed.FS

func TestFileServerEmbed(t *testing.T) {
	dir, err := fs.Sub(assets, "assets")
	if err != nil {
		panic(err)
	}

	mux := new(http.ServeMux)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.FS(dir))))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
