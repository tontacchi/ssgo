package ast

import (
	"fmt"
	"slices"
	"strings"
)

//------------------------------------------------------------------------------
// LeafNode Definition
//------------------------------------------------------------------------------

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


//------------------------------------------------------------------------------
// Public API
//------------------------------------------------------------------------------

func (n LeafNode) ToHTML() (string, error) {
	// guard: leaf must have a value
	if n.Value == "" { return "", fmt.Errorf("leaf node must have a value") }

	// early: no tag -> just the value
	if n.Tag == "" { return n.Value, nil }

	res := fmt.Sprintf("<%s%s>%s</%s>", n.Tag, propsToHTML(n.Props), n.Value, n.Tag)
	return res, nil
}

func (n LeafNode) String() string {
	// guard: no props
	if len(n.Props) == 0 {
		return fmt.Sprintf("LeafNode(%s, %s, map[])", n.Tag, n.Value)
	}

	// sort keys, build sorted map output
	keys := getSortedKeys(n.Props)
	props := []string{}

	for _, key := range keys {
		pair := fmt.Sprintf("%s:%s", key, n.Props[key])
		props = append(props, pair)
	}

	return fmt.Sprintf("LeafNode(%s, %s, map[%s])",
		n.Tag,
		n.Value,
		strings.Join(props, " "),
	)
}


//------------------------------------------------------------------------------
// Helper Methods
//------------------------------------------------------------------------------

func propsToHTML(props map[string]string) string {
	// guard: no key-value pairs to output
	if len(props) == 0 { return "" }

	// hack: sort a key slice for "ordered map" behavior
	keys := make([]string, 0, len(props))
	for key := range props {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	var builder strings.Builder

	// props always come after Tag name, so always has " " before all pairs
	for _, key := range keys {
		fmt.Fprintf(&builder, ` %s="%s"`, key, props[key])
	}

	return builder.String()
}

func getSortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	return keys
}

