package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "world")
	})

	mux.HandleFunc("/image/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "image")
	})

	mux.HandleFunc("/image/picture/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "picture")
	})

	mux.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.Header)
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	fmt.Println("server running")
}
