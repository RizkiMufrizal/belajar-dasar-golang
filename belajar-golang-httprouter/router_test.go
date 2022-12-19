package main

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

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/hello", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello World", string(bytes), "Must Hello World")
}

func TestRouterParam(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Product "+params.ByName("id"))
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/product/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1", string(bytes), "Must Product 1")
}

func TestRouterMultiParam(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/item/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Product "+params.ByName("id")+" dengan item "+params.ByName("itemId"))
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/product/1/item/A1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 dengan item A1", string(bytes), "Must Product 1 dengan item A1")
}

func TestRouterParamCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "gambar "+params.ByName("image"))
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/images/src/icon.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "gambar /src/icon.png", string(bytes), "Must gambar /src/icon.png")
}

//go:embed resources
var resourcesFolder embed.FS

func TestRouterServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resourcesFolder, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:8080/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "hello from file", string(bytes), "Must hello from file")
}

func TestRouterPaanic(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer, "panic : ", error)
	}
	router.GET("/hello", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		panic("ups")
	})
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "panic : ups", string(bytes), "Must panic : ups")
}

func TestRouterNotFound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(404)
		fmt.Fprint(writer, "tidak ketemu")
	})
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, 404, response.StatusCode, "Must 404")
	assert.Equal(t, "tidak ketemu", string(bytes), "Must tidak ketemu")
}

func TestRouterMethodNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(405)
		fmt.Fprint(writer, "methot tidak allow")
	})
	router.POST("/hello", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, 405, response.StatusCode, "Must 405")
	assert.Equal(t, "methot tidak allow", string(bytes), "Must methot tidak allow")
}

type LogMiddleware struct {
	Handler http.Handler
}

func (logMiddleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive Request")
	logMiddleware.Handler.ServeHTTP(writer, request)
}

func TestRouterMiddleware(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(405)
		fmt.Fprint(writer, "methot tidak allow")
	})
	router.GET("/hello", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	middleware := &LogMiddleware{
		Handler: router,
	}

	middleware.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello World", string(bytes), "Must Hello World")
}
