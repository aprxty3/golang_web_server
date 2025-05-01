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

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive Middleware")
	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	middleware := LogMiddleware{router}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorde := httptest.NewRecorder()

	middleware.ServeHTTP(recorde, request)

	response := recorde.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	assert.Equal(t, "Hello World", string(body))
}
