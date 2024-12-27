## Introduction

This library is designed for experimentation and aims to provide a simple way to render emails in Go. Unlike most libraries available for creating HTML emails, this project is inspired by [react-email](https://github.com/resend/react-email) and leverages Go's native HTML templating capabilities to generate clean and reusable email templates.

### Why This Exists

If you're exploring email rendering in Go and finding a lack of flexible, developer-friendly solutions, this library may help. While it is still experimental, it serves as a learning opportunity and a starting point for building feature-rich email templates.

### Key Features

- Leverages Go’s `html/template` for rendering clean, reusable HTML layouts.
- Inspired by React/html concepts, promoting a tree-like hierarchy for building email components.
- **Note:** This package is experimental and should be used for learning or non-critical email systems only.

---

### Example Usage

Here’s a simple example of how to use this library to build and render an email with embedded code formatting:

```go
package main

import (
    "fmt"
    "github.com/sanchitrk/gomailify"
)

func main() {

    // Create a new HTML email template
    html := gomailify.NewHTML("")

    // Sample Go code block to include in the email
    code := `
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World")
    }`

    // Create a code node from the Go code snippet with syntax highlighting
    cb := gomailify.NewCodeNode(code, "go")

    // Add the code block node to the HTML email
    html.AddChild(cb)

    // Render the final HTML email content
    out, err := html.Render()
    if err != nil {
        panic(err)
    }

    fmt.Println(out)
}
```

This will render an email with HTML content including the provided Go code block, formatted correctly and embedded within the body of the email.

---

### Notes

- This library is **experimental** and meant for learning purposes, rapid prototyping, or non-production environments.
- It focuses on simplicity and may lack advanced features like responsive design or in-depth email client compatibility.

Feel free to explore, experiment, and extend!