// Package ast
package ast

import (
	"fmt"
	"strings"
)

type HTMLNode struct {
	Tag      string
	Value    string
	Children []HTMLNode
	Props    map[string]string
}

func NewHTMLNode(tag, value string,
	children []HTMLNode,
	props map[string]string,
) HTMLNode {
	return HTMLNode{
		Tag:      tag,
		Value:    value,
		Children: children,
		Props:    props,
	}
}

func (n HTMLNode) toHTML() (string, error) {
	return "", fmt.Errorf("toHTML not implemented for HTMLNode")
}

func (n HTMLNode) PropsToHTML() string {
	if len(n.Props) == 0 { return "" }

	var builder strings.Builder

	for key, value := range n.Props {
		fmt.Fprintf(&builder, ` %s="%s"`, key, value)
	}

	return builder.String()
}

func (n HTMLNode) String() string {
	return fmt.Sprintf(
		"HTMLNode(%s, %s, %v, %v)",
		n.Tag,
		n.Value,
		n.Children,
		n.Props,
	)
}

