package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World!")
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "www.google.com", http.StatusTemporaryRedirect)
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "www.google.com", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
