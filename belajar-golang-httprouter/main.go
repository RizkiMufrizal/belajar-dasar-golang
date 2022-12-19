package main

import (
	"embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"net/http"
)

//go:embed resources
var resources embed.FS

func main() {
	router := httprouter.New()
	router.GET("/hello", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})
	router.GET("/product/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Product "+params.ByName("id"))
	})
	router.GET("/product/:id/item/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Product "+params.ByName("id")+" dengan item "+params.ByName("itemId"))
	})
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "gambar "+params.ByName("image"))
	})
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	fmt.Println("Server Running")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
