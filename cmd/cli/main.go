package main

import (
	"catcatgo/internal/model"
	"catcatgo/internal/parser"
	"catcatgo/internal/search"
	"catcatgo/internal/storage"
	"fmt"
	"log"
	"os"
)

func Index(path string) []model.Function {
	source, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	
	p := parser.NewParser()
	defer p.Close()

	ast := p.Parse(source, nil)
	defer ast.Close()

	root := ast.RootNode()
	return parser.ExtractFunctions(root, source)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: catcatcli <query>")
		return
	}

	dataPath := "data/storage.json"
	sampleSpace, err := storage.Load(dataPath)
	if err != nil {
		sampleSpace = Index("dump/stb_image.h")
		err := storage.Save(dataPath, sampleSpace)
		if err != nil {
			panic(err)
		}
	}

	query := os.Args[1]
	resultSet := search.Linear(sampleSpace, query)
	for _, function := range resultSet {
		fmt.Println(function.String())
	}
}
