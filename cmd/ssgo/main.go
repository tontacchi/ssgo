package main

import (
	"fmt"

	"ssgo/internal/ast"
)

func main() {
	fmt.Println("ssgo: started")

	node := ast.NewLeafNode("a", "click", map[string]string{
		"href": "https://example.com",
	})
	fmt.Println(node.String())

	fmt.Println("ssgo: finished")
}
