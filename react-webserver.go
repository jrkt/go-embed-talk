package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed react-app/build
var embeddedFiles embed.FS

func main() {
	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	buildFS, err := fs.Sub(embeddedFiles, "react-app/build")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(buildFS)))

	log.Println("serving on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
