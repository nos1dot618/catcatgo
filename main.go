package main

import (
	"fmt"
	"os"
	"log"

	ts "github.com/tree-sitter/go-tree-sitter"
	tsc "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

func getText(node *ts.Node, source []byte) string {
	return string(source[node.StartByte(): node.EndByte()])
}

func printFunction(node *ts.Node, source []byte) {
	declarator := node.ChildByFieldName("declarator")
	if declarator == nil {
		// TODO: Maybe log this.
		return
	}

	nameNode := declarator.ChildByFieldName("declarator")
	params := declarator.ChildByFieldName("parameters")

	functionName := getText(nameNode, source)
	paramTypes := []string{}

	if params != nil {
		for i := uint(0); i < params.NamedChildCount(); i++ {
			param := params.NamedChild(i)
			if param.Kind() == "parameter_declaration" {
				paramTypeNode := param.ChildByFieldName("type")
				if paramTypeNode != nil {
					paramTypes = append(paramTypes, getText(paramTypeNode, source))
				}
			}
		}
	}

	returnTypeNode := node.ChildByFieldName("type")
	returnType := "void"
	if returnTypeNode != nil {
		returnType = getText(returnTypeNode, source)
	}

	// Print function signature in style of Haskell.
	fmt.Printf("%s :: ", functionName)
	for _, paramType := range paramTypes {
		fmt.Printf("%s -> ", paramType)
	}
	fmt.Printf("%s\n", returnType)
}

func walk(node *ts.Node, source []byte) {
	if node.Kind() == "function_definition" {
		printFunction(node, source)
	}
	
	for i := uint(0); i < node.ChildCount(); i++ {
		child := node.Child(i)
		walk(child, source)
	}
}

func main() {
	source, err := os.ReadFile("dump/stb_image.h")
	if err != nil {
		log.Fatal(err)
	}
	
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(ts.NewLanguage(tsc.Language()))

	ast := parser.Parse(source, nil)
	defer ast.Close()

	root := ast.RootNode()
	walk(root, source)
}
