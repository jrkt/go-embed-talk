// embed_fs.go
package main

import (
	"embed"
)

//go:embed templates
var f embed.FS

func main() {
	data, _ := f.ReadFile("templates/hello.txt")
	print(string(data))
}
