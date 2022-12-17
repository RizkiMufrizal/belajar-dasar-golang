package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	fmt.Println(response.Status)
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func QueryParameterHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "hello")
	} else {
		fmt.Fprint(writer, "hello "+name)
	}
}

func TestQueryParameter(t *testing.T) {
	requestWithParameter := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=rizki", nil)
	recorderWithParameter := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	QueryParameterHandler(recorderWithParameter, requestWithParameter)
	QueryParameterHandler(recorder, request)

	responseWithParameter := recorderWithParameter.Result()
	response := recorder.Result()

	bodyWithParameter, _ := io.ReadAll(responseWithParameter.Body)
	bodyStringWithParameter := string(bodyWithParameter)

	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, "hello rizki", bodyStringWithParameter, "Must hello rizki")
	assert.Equal(t, "hello", bodyString, "Must hello")
}

func QueryMultiParameterHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	fmt.Fprint(writer, "hello "+strings.Join(query["name"], ","))
}

func TestQueryMultiParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=rizki&name=mufrizal", nil)
	recorder := httptest.NewRecorder()
	QueryMultiParameterHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, "hello rizki,mufrizal", bodyString, "Must hello rizki,mufrizal")
}
