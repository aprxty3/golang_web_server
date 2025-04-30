package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleWare *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Execute before handle")
	middleWare.Handler.ServeHTTP(writer, request)
	fmt.Println("Execute after handle")
}

func TestMiddlerware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Execute")
		fmt.Fprint(writer, "Hello World Middleware")
	})

	logMiddleware := &LogMiddleware{Handler: mux}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: logMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
