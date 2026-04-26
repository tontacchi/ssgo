package ast

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
	res := ""

	return res, nil
}
