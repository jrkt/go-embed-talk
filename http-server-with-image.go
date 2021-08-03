// http-server-with-image.go
package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed templates
var content embed.FS

func main() {
	mux := http.NewServeMux()
	mux.Handle("/templates/assets/", http.FileServer(http.FS(content)))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index, _ := content.ReadFile("templates/http-with-image-index.html")
		w.Write(index)
	})

	log.Println("serving on :8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
