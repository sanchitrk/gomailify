package gomailify

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"html/template"
	"strings"
)

// highlightCode generates syntax-highlighted HTML for the given code and language.
func highlightCode(code, language string) (string, error) {
	// Get lexer for the specified language
	lexer := lexers.Get(language)
	if lexer == nil {
		// Try to guess the lexer if language not specified
		lexer = lexers.Analyse(code)
		if lexer == nil {
			lexer = lexers.Fallback
		}
	}

	// Use the "monokai" style for highlighting
	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}

	// Create HTML formatter with line numbers and inline styles
	formatter := html.New(
		html.WithLineNumbers(true),
		html.LineNumbersInTable(true),
		html.WithClasses(false), // Inline styles instead of CSS classes
		html.TabWidth(4),
	)

	// Tokenize the code
	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return "", fmt.Errorf("tokenization error: %v", err)
	}

	// Generate HTML with inline styles
	var sb strings.Builder
	err = formatter.Format(&sb, style, iterator)
	if err != nil {
		return "", fmt.Errorf("formatting error: %v", err)
	}

	// Wrap the output in a styled <div>
	//result := fmt.Sprintf(`
	//	<div style="background-color: #272822; padding: 10px; border-radius: 5px; font-family: 'Courier New',monospace,Consolas,Monaco,'Andale Mono';">
	//	%s
	//	</div>`, sb.String())

	return sb.String(), nil
}

type Node interface {
	Render() (string, error)
	GetChildren() []Node
	AddChild(node Node)
}

type BaseNode struct {
	children []Node
	value    string
	tmpl     *template.Template
}

func (b *BaseNode) GetChildren() []Node {
	return b.children
}

func (b *BaseNode) AddChild(node Node) {
	b.children = append(b.children, node)
}

type NodeData struct {
	Value    template.HTML
	Children []template.HTML
}

type HTML struct {
	BaseNode
}

func NewHTML(value string) *HTML {
	tmpl := template.Must(template.New("html").Parse(HTMLTemplate))

	return &HTML{
		BaseNode: BaseNode{
			value:    value,
			children: make([]Node, 0),
			tmpl:     tmpl,
		},
	}
}

func (h *HTML) Render() (string, error) {
	// First, render all children
	childrenOutput := make([]template.HTML, 0, len(h.children))
	for _, child := range h.children {
		rendered, err := child.Render()
		if err != nil {
			return "", err
		}
		childrenOutput = append(childrenOutput, template.HTML(rendered))
	}

	// Prepare the data for template
	data := NodeData{
		Value:    template.HTML(h.value),
		Children: childrenOutput,
	}

	// Execute the template
	var buf bytes.Buffer
	if err := h.tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type Paragraph struct {
	BaseNode
}

func NewParagraph(value string) *Paragraph {
	fmt.Println("** Creating paragraph **")
	// Create the template for paragraph elements
	tmpl := template.Must(template.New("paragraph").Parse(ParagraphTemplate))

	return &Paragraph{
		BaseNode: BaseNode{
			value:    value,
			children: make([]Node, 0),
			tmpl:     tmpl,
		},
	}
}

func (p *Paragraph) Render() (string, error) {
	// First, render all children
	childrenOutput := make([]template.HTML, 0, len(p.children))
	for _, child := range p.children {
		rendered, err := child.Render()
		if err != nil {
			return "", err
		}
		childrenOutput = append(childrenOutput, template.HTML(rendered))
	}

	// Prepare the data for template
	data := NodeData{
		Value:    template.HTML(p.value),
		Children: childrenOutput,
	}

	// Execute the template
	var buf bytes.Buffer
	if err := p.tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type Container struct {
	BaseNode
}

func NewContainer(value string) *Container {
	// Create the template for paragraph elements
	tmpl := template.Must(template.New("container").Parse(ContainerTemplate))

	return &Container{
		BaseNode: BaseNode{
			value:    value,
			children: make([]Node, 0),
			tmpl:     tmpl,
		},
	}
}

func (c *Container) Render() (string, error) {
	// First, render all children
	childrenOutput := make([]template.HTML, 0, len(c.children))
	for _, child := range c.children {
		rendered, err := child.Render()
		if err != nil {
			return "", err
		}
		childrenOutput = append(childrenOutput, template.HTML(rendered))
	}

	// Prepare the data for template
	data := NodeData{
		Value:    template.HTML(c.value),
		Children: childrenOutput,
	}

	// Execute the template
	var buf bytes.Buffer
	if err := c.tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type CodeNode struct {
	BaseNode
	code     string
	language string
}

func NewCodeNode(code, language string) *CodeNode {
	tmpl := template.Must(template.New("code").Parse(CodeBlockTemplate))

	return &CodeNode{
		BaseNode: BaseNode{
			children: make([]Node, 0),
			tmpl:     tmpl,
		},
		code:     code,
		language: language,
	}
}

func (c *CodeNode) Render() (string, error) {
	highlighted, err := highlightCode(c.code, c.language)
	if err != nil {
		return "", fmt.Errorf("error highlighting code: %v", err)
	}

	childrenOutput := make([]template.HTML, 0, len(c.children))
	for _, child := range c.children {
		rendered, err := child.Render()
		if err != nil {
			return "", err
		}
		childrenOutput = append(childrenOutput, template.HTML(rendered))
	}

	data := NodeData{
		Value:    template.HTML(highlighted),
		Children: childrenOutput,
	}

	var buf bytes.Buffer
	if err := c.tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type TextNode struct {
	BaseNode
}

func NewTextNode(value string) *TextNode {
	tmpl := template.Must(template.New("text").Parse(TextTemplate))
	return &TextNode{
		BaseNode: BaseNode{
			value:    value,
			children: make([]Node, 0),
			tmpl:     tmpl,
		},
	}
}

func (t *TextNode) Render() (string, error) {
	childrenOutput := make([]template.HTML, 0, len(t.children))
	for _, child := range t.children {
		rendered, err := child.Render()
		if err != nil {
			return "", err
		}
		childrenOutput = append(childrenOutput, template.HTML(rendered))
	}

	// Prepare the data for template
	data := NodeData{
		Value:    template.HTML(t.value),
		Children: childrenOutput,
	}

	// Execute the template
	var buf bytes.Buffer
	if err := t.tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
