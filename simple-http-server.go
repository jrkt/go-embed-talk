// simple-http-server.go
package main

import (
	_ "embed"
	"log"
	"net/http"
)

//go:embed templates/simple-http-index.html
var indexPage []byte

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(indexPage)
	})

	log.Println("serving on :8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
