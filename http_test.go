package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GGWP"))
	})
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This my index"))
}

func TestHandlerHttp(t *testing.T) {
	http.HandleFunc("/index", IndexHandler)

	http.ListenAndServe(":8080", nil)
}

func TestServeMux(t *testing.T) {
	mux := new(http.ServeMux)

	mux.Handle("/index/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hophop"))
	}))

	mux.HandleFunc("/index/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("dashboard"))
	})

	server := new(http.Server)

	server.Addr = ":8080"
	server.Handler = mux
	server.ListenAndServe()
}

func TestRequest(t *testing.T) {
	http.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, r.Method)
			fmt.Fprintln(w, r.RequestURI)
			fmt.Fprintln(w, r.Header)
			fmt.Fprintln(w, r.Host)
		},
	))
	http.ListenAndServe(":8080", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloWorld(recorder, request)

	result := recorder.Result()
	resultString, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success:", string(resultString))
	assert.Equal(t, "Hello World", string(resultString))
}

func MultiplerParam(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	address := r.URL.Query().Get("address")

	fmt.Fprintf(w, "I am %s. I'm from %s", name, address)
}

func TestMultiParam(t *testing.T) {
	name := "Richard"
	address := "USA"
	url := "http://localhost:8080/inctoduce?name=" + name + "&address=" + address
	request := httptest.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()

	MultiplerParam(response, request)

	result, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	stringResult := string(result)
	assert.Equal(t, "I am "+name+"."+" I'm from "+address, stringResult)
	fmt.Println(stringResult)

}

func MultiSameParam(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	address := queryParam["address"]
	fmt.Fprint(w, strings.Join(address, " "))
}

func TestMutilSameParam(t *testing.T) {
	address1 := "USA"
	address2 := "India"
	url := "http://localhost:8080/sameparam?address=" + address1 + "&address=" + address2
	request := httptest.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()

	MultiSameParam(response, request)

	result, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	stringResult := string(result)
	fmt.Println(stringResult)

	assert.Equal(t, address1+" "+address2, stringResult)
}

func GetHeaderType(w http.ResponseWriter, r *http.Request) {
	headerType := r.Header.Get("content-type")
	fmt.Fprint(w, headerType)
}

func TestHeaderType(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")

	response := httptest.NewRecorder()

	GetHeaderType(response, request)
	result := response.Result()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Liky-Dev")
	fmt.Fprint(w, "Success")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("POST", "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	response := httptest.NewRecorder()

	ResponseHeader(response, request)
	result := response.Result()

	data, _ := io.ReadAll(result.Body)

	fmt.Println(string(data))
	fmt.Println(result.Header.Get("X-Powered-By"))
}
