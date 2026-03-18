package parser

import (
	ts "github.com/tree-sitter/go-tree-sitter"
	tsc "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

func NewParser() *ts.Parser {
	parser := ts.NewParser()
	parser.SetLanguage(ts.NewLanguage(tsc.Language()))
	return parser;
}
