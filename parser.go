package gomailify

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"html/template"
	"strings"
)

// CodeToHTML generates the HTML string from the inline code.
func CodeToHTML(code string) (string, error) {
	// Parse the template.
	tmpl, err := template.New("codeToHTML").Parse(InlineCodeTemplate)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %w", err)
	}

	// Execute the template.
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, code); err != nil {
		return "", fmt.Errorf("error executing template: %w", err)
	}

	// Return the generated HTML string.
	return buf.String(), nil
}

// HTMLFragmentToNodes parses an HTML fragment string with a given context and returns a slice of nodes.
func HTMLFragmentToNodes(fragmentHTML string, context *html.Node) ([]*html.Node, error) {
	dumbContext := &html.Node{
		Type:     context.Type,
		Data:     context.Data,
		DataAtom: context.DataAtom,
	}
	// Parse the HTML fragment using the provided context node.
	nodes, err := html.ParseFragment(strings.NewReader(fragmentHTML), dumbContext)
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML fragment: %w", err)
	}
	return nodes, nil
}

// ReplaceChildren TODO: FIX ME
func ReplaceChildren(parentNode *html.Node, newChildren []*html.Node) {
	// Remove all existing children
	for child := parentNode.FirstChild; child != nil; {
		next := child.NextSibling
		parentNode.RemoveChild(child)
		child = next
	}

	// Append new children
	for _, child := range newChildren {
		// Detach the node from its current parent, if necessary
		if child.Parent != nil {
			child.Parent.RemoveChild(child)
		}

		// Append the node to parentNode
		parentNode.AppendChild(child)
	}
}

func extractParagraphAttributes(templateStr string) (map[string]string, error) {
	attributes := make(map[string]string)
	doc, err := html.Parse(strings.NewReader(templateStr))
	if err != nil {
		return nil, fmt.Errorf("invalid HTML template: %v", err)
	}

	// Find our p tag
	var findP func(*html.Node) *html.Node
	findP = func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "p" {
			return n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if found := findP(c); found != nil {
				return found
			}
		}
		return nil
	}

	pNode := findP(doc)
	if pNode == nil {
		return nil, fmt.Errorf("no p tag found in template")
	}

	// Extract attributes
	for _, attr := range pNode.Attr {
		attributes[attr.Key] = attr.Val
	}

	return attributes, nil
}

func InlineCodeFromNode(n *html.Node) string {
	var code string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			return code + c.Data
		}
	}
	return code
}

func DepthFirstTraversal(n *html.Node, depth int) {
	switch n.Type {
	case html.ElementNode:
		switch n.Data {
		case "p":
			newAttr := make([]html.Attribute, 0)
			attributes, err := extractParagraphAttributes(ParagraphTemplate)
			if err != nil {
				panic(err)
			}
			for k, v := range attributes {
				newAttr = append(newAttr, html.Attribute{
					Key: k,
					Val: v,
				})
			}
			n.Attr = newAttr
		case "code":
			// check if its inline or code block
			var isInline bool
			var parentNode *html.Node
			if n.Parent != nil && n.Parent.Data == "p" {
				isInline = true
				parentNode = n.Parent
			}
			if isInline {
				code := InlineCodeFromNode(n)
				codeHTML, err := CodeToHTML(code)
				if err != nil {
					panic(err)
				}
				nodes, err := HTMLFragmentToNodes(codeHTML, parentNode)
				if err != nil {
					panic(err)
				}
				ReplaceChildren(parentNode, nodes)
			}
		}
	case html.TextNode:
		break
	default:
		break
	}

	// Recursively traverse child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		DepthFirstTraversal(c, depth+1)
	}
}

func Render(doc *html.Node) string {
	var buf bytes.Buffer
	err := html.Render(&buf, doc)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
