package ast

import "testing"

func TestLeafToHTMLParagraph(t *testing.T) {
	node := NewLeafNode("p", "Hello, world!", nil)

	got, err := node.ToHTML()
	if err != nil { t.Fatalf("Unexpected error: %v", err) }

	want := "<p>Hello, world!</p>"
	if got != want { t.Fatalf("got %q, want %q", got, want) }
}

func TestLeafToHTMLAnchorWithProps(t *testing.T) {
	node := NewLeafNode(
		"a", 
		"Click me!",
		map[string]string{"href": "https://www.google.com"},
	)
	
	got, err := node.ToHTML()
	if err != nil { t.Fatalf("Unexpected error: %v", err) }

	want := `<a href="https://www.google.com">Click me!</a>`
	if got != want { t.Fatalf("got %q, want %q", got, want) }
} 

