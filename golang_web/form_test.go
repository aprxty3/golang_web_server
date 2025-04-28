package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func PostForm(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.Form.Get("first_name")
	lastName := request.Form.Get("last_name")

	fmt.Fprintf(writer, "Hello %s, %s", firstName, lastName)
}

func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("first_name=John&last_name=Doe")
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	PostForm(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
