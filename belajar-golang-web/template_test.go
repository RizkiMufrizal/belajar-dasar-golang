package belajargolangweb

import (
	"embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func SimpleTempletData(writer http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Sample Title",
		"Name":  "Rizki",
		"Address": map[string]interface{}{
			"Street": "Akses UI",
		},
	})
}

func TestSimpleTempletData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleTempletData(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

type TemplateSampleStruct struct {
	Title   string
	Name    string
	Address Address
	Nilai   int32
}

type Address struct {
	Street string
}

func SimpleTempletDataStruct(writer http.ResponseWriter, r *http.Request, action *TemplateSampleStruct) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "name.gohtml", action)
}

func TestSimpleTempletDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	action := TemplateSampleStruct{
		Title: "Sample Title",
		Name:  "Rizki",
		Nilai: 90,
		Address: Address{
			Street: "Akses UI",
		},
	}

	SimpleTempletDataStruct(recorder, request, &action)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func SimpleTempleteAction(writer http.ResponseWriter, r *http.Request, action *TemplateSampleStruct) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "templateaction.gohtml", action)
}

func TestSimpleTempletAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	action := TemplateSampleStruct{
		Title: "Sample Title",
		Name:  "Rizki",
		Nilai: 90,
		Address: Address{
			Street: "Akses UI",
		},
	}

	SimpleTempleteAction(recorder, request, &action)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func SimpleTempleteArray(writer http.ResponseWriter, r *http.Request, action *map[string]interface{}) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "templateaction.gohtml", action)
}

func TestSimpleTempleteArray(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	action := map[string]interface{}{
		"Hobbies": []map[string]interface{}{
			{
				"Olahraga": "Mtb",
			},
			{
				"Olahraga": "Bola",
			},
		},
	}

	SimpleTempleteArray(recorder, request, &action)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestSimpleTempleteWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	action := TemplateSampleStruct{
		Title: "Sample Title",
		Name:  "Rizki",
		Nilai: 90,
		Address: Address{
			Street: "Akses UI",
		},
	}

	SimpleTempletDataStruct(recorder, request, &action)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func SimpleTempleteLayout(writer http.ResponseWriter, r *http.Request, action *map[string]interface{}) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "layout", action)
}

func TestSimpleTempleteLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	action := map[string]interface{}{
		"Name":  "Rizki",
		"Title": "Belajar Golang",
	}

	SimpleTempleteLayout(recorder, request, &action)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

type MyPage struct {
	Name string
}

func (page MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + page.Name
}

func SimpleTempleteFuction(writer http.ResponseWriter, r *http.Request, action *MyPage) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "sample", action)
}

func TestSimpleFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	action := MyPage{
		Name: "Rizki",
	}

	SimpleTempleteFuction(recorder, request, &action)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func SimpleTempleteCreateGlobalFuction(writer http.ResponseWriter, r *http.Request, action *MyPage) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	tg, err := t.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	tg.ExecuteTemplate(writer, "sample", action)
}

func TestSimpleCreateGlobalFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	action := MyPage{
		Name: "Rizki",
	}

	SimpleTempleteCreateGlobalFuction(recorder, request, &action)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

var t = template.New("FUNCTION").Funcs(map[string]interface{}{
	"upper": func(value string) string {
		return strings.ToUpper(value)
	},
})
var tg, _ = t.ParseFS(templates, "templates/*.gohtml")

func SimpleTempleteCaching(writer http.ResponseWriter, r *http.Request, action *MyPage) {
	tg.ExecuteTemplate(writer, "sample", action)
}

func TestSimpleCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	action := MyPage{
		Name: "Rizki",
	}

	SimpleTempleteCaching(recorder, request, &action)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func SimpleTempleteXSS(writer http.ResponseWriter, r *http.Request) {
	tg.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Go-lang",
		"Body":  template.HTML("<p>Selamat Belajar Golang</p>"),
	})
}

func TestSimpleXss(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleTempleteXSS(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
