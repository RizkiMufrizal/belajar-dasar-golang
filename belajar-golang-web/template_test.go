package belajargolangweb

import (
	"embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "SIMPLE", "Hello From Template")
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "<html><body>Hello From Template</body></html>", string(body), "Must <html><body>Hello From Template</body></html>")
}

func SimpleFileHtml(writer http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello File Template")
}

func TestSimpleFileHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleFileHtml(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func SimpleFileDirectoryHtml(writer http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello File Template")
}

func TestSimpleFileDirectoryHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleFileDirectoryHtml(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func SimpleFileDirectoryEmbedHtml(writer http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello File Template")
}

func TestSimpleFileDirectoryEmbedHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleFileDirectoryEmbedHtml(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
