// embed_byte_slice.go
package main

import _ "embed"

//go:embed templates/hello.txt
var b []byte

//go:embed templates/hello.txt
var s string

func main() {
	print("From byte slice: ", string(b), "\n")
	print("From string: ", s)
}
