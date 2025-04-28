package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "SetCookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprintln(writer, err)
	} else {
		fmt.Fprintln(writer, "Hello",
			cookie.Value)
	}

}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set", SetCookie)
	mux.HandleFunc("/get", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/?name=Aji", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	response := recorder.Result().Cookies()

	for _, cookie := range response {
		fmt.Printf("Cookie name: %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Aji"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
