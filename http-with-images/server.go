// server.go
package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed index-with-images.html
var indexPage []byte

//go:embed assets
var assets embed.FS

func main() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.FileServer(http.FS(assets)))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(indexPage)
	})
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
