package go_router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Gak Boleh")
	})

	//router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//	fmt.Fprint(writer, "Get")
	//})

	router.POST("/", func(w http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Post")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorde := httptest.NewRecorder()

	router.ServeHTTP(recorde, request)

	response := recorde.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	assert.Equal(t, "Gak Boleh", string(body))
}
