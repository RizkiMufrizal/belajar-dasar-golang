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

func HeaderHandler(writer http.ResponseWriter, request *http.Request) {
	header := request.Header.Get("Content-Type")
	writer.Header().Add("X-CUSTOM", "belajar golang")
	fmt.Fprint(writer, "header "+header)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	HeaderHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, "header application/json", bodyString, "Must header")
	assert.Equal(t, "belajar golang", response.Header.Get("X-CUSTOM"), "Must belajar golang")
}

func FormPostHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	firstName := request.PostForm.Get("firstName")
	lastName := request.PostForm.Get("lastName")
	fmt.Fprintf(writer, "hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstName=rizki&lastName=mufrizal")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/post", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	FormPostHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, "hello rizki mufrizal", bodyString, "Must rizki mufrizal")
}

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(400)
		fmt.Fprint(writer, "Bad Request")
	} else {
		writer.WriteHeader(200)
		fmt.Fprint(writer, "Ok")
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=rizki", nil)

	recorder := httptest.NewRecorder()
	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, "Ok", bodyString, "Must Ok")
	assert.Equal(t, 200, response.StatusCode, "Must 200")
}

func TestResponseCodeBadRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)

	recorder := httptest.NewRecorder()
	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, "Bad Request", bodyString, "Must Bad Request")
	assert.Equal(t, 400, response.StatusCode, "Must 400")
}
