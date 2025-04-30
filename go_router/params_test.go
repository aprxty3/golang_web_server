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

func TestRouterParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Product" + params.ByName("id")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/123", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)

	fmt.Println(string(byte))

	assert.Equal(t, "Product123", string(byte))
}
