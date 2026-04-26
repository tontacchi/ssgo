package ast

import "testing"

func TestPropsToHTML(t *testing.T) {
	node := NewHTMLNode(
		"a",
		"click me", 
		nil, 
		map[string]string{
			"href":   "https://www.google.com",
			"target": "_blank",
		},
	)

	actual := node.PropsToHTML()

	// Map order is not guaranteed in Go, so check both valid outputs.
	want1 := ` href="https://www.google.com" target="_blank"`
	want2 := ` target="_blank" href="https://www.google.com"`

	if actual != want1 && actual != want2 {
		t.Fatalf("got %q, want %q or %q", actual, want1, want2)
	}
}

func TestPropsToHTMLEmpty(t *testing.T) {
	node := NewHTMLNode("p", "Hello", nil, nil)

	expected, actual := "", node.PropsToHTML()

	if actual != expected {
		t.Fatalf("actual=%q, expected=%q", actual, expected)
	}
}

