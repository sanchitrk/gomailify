package main

import (
	"fmt"
	"github.com/sanchitrk/gomailify"
)

func main() {

	html := gomailify.NewHTML("")

	code := `
	package main
	
	import "fmt"
	
	func main() {
		fmt.Println("Hello, World")
	}`

	cb := gomailify.NewCodeNode(code, "go")
	html.AddChild(cb)

	out, err := html.Render()
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
