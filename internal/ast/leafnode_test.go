package ast

import "testing"


//------------------------------------------------------------------------------
// Positive Tests
//------------------------------------------------------------------------------

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

func TestLeafToHTMLRawText(t *testing.T) {
	node := NewLeafNode("", "raw text", nil)

	got, err := node.ToHTML()
	if err != nil { t.Fatalf("Unexpected error: %v", err) }

	want := "raw text"
	if got != want { t.Fatalf("got %q, want %q", got, want) }
}

func TestLeafToHTMLCode(t *testing.T) {
	node := NewLeafNode("code", `fmt.Println("hi")`, nil)

	got, err := node.ToHTML()
	if err != nil { t.Fatalf("Unexpected error: %v", err) }

	want := `<code>fmt.Println("hi")</code>`	
	if got != want { t.Fatalf("got %q, want %q", got, want) }
}

func TestLeafToHTMLImage(t *testing.T) {
	node := NewLeafNode("img", " ", map[string]string{
		"src": "image.png",
		"alt": "example image",
	})

	got, err := node.ToHTML()
	if err != nil { t.Fatalf("Unexpected error: %v", err) }

	// reminder: props are sorted by keys lexographically
	want := `<img alt="example image" src="image.png"> </img>`	
	if got != want { t.Fatalf("got %q, want %q", got, want) }
	
}


//------------------------------------------------------------------------------
// Negative Tests
//------------------------------------------------------------------------------

func TestLeafToHTMLErrorsWithoutValue(t *testing.T) {
	node := NewLeafNode("p", "", nil)

	_, err := node.ToHTML()
	if err == nil {
		t.Fatalf("expected error for empty value")
	}
}

