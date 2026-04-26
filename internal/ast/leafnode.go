package ast

import "fmt"

type LeafNode struct {
	Tag   string
	Value string
	Props map[string]string
}

func NewLeafNode(tag, value string, props map[string]string) LeafNode {
	return LeafNode{
		Tag:   tag,
		Value: value,
		Props: props,
	}
}

func (n LeafNode) ToHTML() (string, error) {
	// guard: leaf must have a value
	if n.Value == "" { return "", fmt.Errorf("leaf node must have a value") }

	// early: no tag -> just the value
	if n.Tag == "" { return n.Value, nil }

	res := fmt.Sprintf("<%s%s>%s<%s>", n.Tag, propsToHTML(n.Props), n.Value, n.Tag)
	return res, nil
}


// helper functions

func propsToHTML(props map[string]string) string {
	return ""
}

