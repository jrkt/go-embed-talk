// embed_string.go
package main

import _ "embed"

//go:embed templates/hello.txt
var s string

func main() {
	print(s)
}
