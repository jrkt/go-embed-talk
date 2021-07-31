package main

import (
	"embed"
	"fmt"
	"io/ioutil"
)

//go:embed test-templates
var f embed.FS

func main() {
	b, err := ioutil.ReadFile("/tmp/test-templates/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	//fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
	//	fmt.Println(path)
	//	return nil
	//})
}
