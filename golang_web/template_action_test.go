package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templateAction embed.FS

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templateAction, "templates/*.gohtml"))

	t.ExecuteTemplate(writer, "action.gohtml", Page{
		Title: "Hello World!",
		Name:  "Gopher",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateActionIf(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
