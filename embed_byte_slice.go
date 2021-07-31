// test.go
package main

import _ "embed"

//go:embed templates/hello.txt
var b []byte

func main() {
	print(string(b))
}
