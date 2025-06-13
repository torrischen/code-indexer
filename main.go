package main

import (
	"fmt"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_go "github.com/tree-sitter/tree-sitter-go/bindings/go"
)

func main() {
	code := []byte("var a = 1")

	parser := tree_sitter.NewParser()
	defer parser.Close()
	parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_go.Language()))

	tree := parser.Parse(code, nil)
	defer tree.Close()

	root := tree.RootNode()
	fmt.Println(root.ToSexp())
}
