package main

import (
	"fmt"
	"github.com/sanchitrk/gomailify"
	"golang.org/x/net/html"
	"strings"
)

func main() {
	htmlDoc := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Sample Page</title>
	</head>
	<body>
		<div class="container">
			<h1>Welcome</h1>
			<p>This is a <strong>sample</strong> paragraph.</p>
			<p>Check the inline code:<code>print("hello,world")</code>Looks like it works!</p>
			<ul>
				<li>Item 1</li>
				<li>Item 2</li>
			</ul>
		</div>
	</body>
	</html>`

	doc, err := html.Parse(strings.NewReader(htmlDoc))
	if err != nil {
		panic(err)
	}

	gomailify.DepthFirstTraversal(doc, 0)
	out := gomailify.Render(doc)
	fmt.Println(out)
}
