package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func Download(writer http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		writer.WriteHeader(400)
		fmt.Fprint(writer, "Bad Request")
		return
	}
	writer.Header().Set("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(writer, r, "./resources/"+file)
}

func TestDownload(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/download", Download)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
