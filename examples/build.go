package main

import (
	"fmt"
	"github.com/sanchitrk/gomailify"
)

func main() {

	html := gomailify.NewHTML("")
	code := `fmt.Println("Hello, World")`
	//code := `
	//package main
	//
	//import "fmt"
	//
	//func main() {
	//	fmt.Println("Hello, World")
	//}`

	cb := gomailify.NewCodeNode(code, "go")
	html.AddChild(cb)
	out, err := html.Render()
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

	//p := gomailify.NewParagraph("Hello, World")
	//out, err := p.Render()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(out)

	//html := gomailify.NewHTML("")
	//p := gomailify.NewParagraph("Hello, World")
	//html.AddChild(p)
	//out, err := html.Render()
	//if err != nil {
	//	panic(err)
	//}

	//text := gomailify.NewTextNode("Hello, World")
	//
	//text.AddChild(gomailify.NewTextNode("Hello, World"))
	//
	//out, err := text.Render()
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println(out)
}
