// Package ast
package ast

import "fmt"

type TextType string
const (
	TextTypeText   = "text"
	TextTypeBold   = "bold"
	TextTypeItalic = "italic"
	TextTypeCode   = "code"
	TextTypeLink   = "link"
	TextTypeImage  = "image"
)

type TextNode struct {
	Text     string
	TextType TextType
	URL      string
}

func NewTextNode(text string, textType TextType, url ...string) TextNode {
	node := TextNode{
		Text:     text,
		TextType: textType,
		URL:      "",
	}

	if len(url) > 0 { node.URL = url[0] }

	return node
}

func (t TextNode) Equal(other TextNode) bool {
	return t.Text == other.Text &&
		t.TextType == other.TextType &&
		t.URL == other.URL
}

func (t TextNode) String() string {
	return fmt.Sprintf("TextNode(%s, %s, %s)", t.Text, t.TextType, t.URL)
}

