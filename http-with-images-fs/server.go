package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed templates
var site embed.FS

func main() {
	mux := http.NewServeMux()
	mux.Handle("/templates/assets/", http.FileServer(http.FS(site)))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index, _ := site.ReadFile("templates/index.html")
		w.Write(index)
	})
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}