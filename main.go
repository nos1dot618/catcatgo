package main

import (
	"catcatgo/internal/parser"
	"fmt"
	"log"
	"os"
)

func main() {
	source, err := os.ReadFile("dump/stb_image.h")
	if err != nil {
		log.Fatal(err)
	}
	
	p := parser.NewParser()
	defer p.Close()

	ast := p.Parse(source, nil)
	defer ast.Close()

	root := ast.RootNode()
	functions := parser.ExtractFunctions(root, source)
	for _, function := range functions {
		fmt.Println(function.String())
	}
}
