// Package ast
package ast

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



