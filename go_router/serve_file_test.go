package go_router

import (
	"embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	director, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(director))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/files/file.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)

	fmt.Println(string(byte))

	assert.Equal(t, "Lesgoo", string(byte))
}
