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

func TestPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + id + " Items " + itemId
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/123/items/123", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)

	fmt.Println(string(byte))

	assert.Equal(t, "Product 123 Items 123", string(byte))
}

func TestPatternCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Images : " + image
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/images/small/picture.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)

	fmt.Println(string(byte))

	assert.Equal(t, "Images : /small/picture.png", string(byte))
}
