package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	fmt.Fprint(writer, contentType)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Golang Web")
	fmt.Fprint(writer, "Hello World")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	request.Header.Add("x-powered-by", "application/json")

	recorder := httptest.NewRecorder()
	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	fmt.Println(response.Header.Get("x-powered-by"))
}
