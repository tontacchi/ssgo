package ast

import "testing"

func TestTextNodeEqual(t *testing.T) {
	node1 := NewTextNode("This is a text node", TextTypeBold)
	node2 := NewTextNode("This is a text node", TextTypeBold)

	if !node1.Equal(node2) {
		t.Fatalf("expected nodes to be equal")
	}
}

func TestTextNodeNotEqualTextType(t *testing.T) {
	node1 := NewTextNode("This is a text node", TextTypeBold)
	node2 := NewTextNode("This is a text node", TextTypeItalic)

	if node1.Equal(node2) {
		t.Fatalf("expected nodes to not be equal")
	}
}

func TestTextNodeWithURL(t *testing.T) {
	node := NewTextNode(
		"This is some anchor text",
		TextTypeLink,
		"https://www.boot.dev",
	)

	want := "TextNode(This is some anchor text, link, https://www.boot.dev)"

	if node.String() != want {
		t.Fatalf("got %q, want %q", node.String(), want)
	}
}
