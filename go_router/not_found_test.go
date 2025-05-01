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

func TestNotFound(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Gak Ketemu")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorde := httptest.NewRecorder()

	router.ServeHTTP(recorde, request)

	response := recorde.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	assert.Equal(t, "Gak Ketemu", string(body))
}
