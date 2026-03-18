package parser

import (
	"catcatgo/internal/model"

	ts "github.com/tree-sitter/go-tree-sitter"
)

func ExtractFunctions(root *ts.Node, source []byte) []model.Function {
	var functions []model.Function
	var walk func(node *ts.Node)

	walk = func(node *ts.Node) {
		if node.Kind() == "function_definition" {
			if function := extractFunction(node, source); function != nil {
				functions = append(functions, *function)
			}
		}

		for i := uint(0); i < node.ChildCount(); i++ {
			walk(node.Child(i))
		}
	}

	walk(root)
	return functions
}

func NormalizeType(s string) string {
	// TODO: Complete this.
	return s
}

func extractFunction(node *ts.Node, source []byte) *model.Function {
	functionDeclaratorNode := node.ChildByFieldName("declarator")
	if functionDeclaratorNode == nil {
		return nil
	}

	functionNameNode := extractIdentifier(functionDeclaratorNode)
	if functionNameNode == nil {
		return nil
	}

	functionName := extractText(functionNameNode, source)
	parameterTypeList := extractFunctionParameters(functionDeclaratorNode, source)
	returnType := extractFunctionReturnType(node, source)

	return &model.Function{
		Name:       functionName,
		Parameters: parameterTypeList,
		ReturnType: returnType,
	}
}

func extractFunctionParameters(functionDeclaratorNode *ts.Node, source []byte) []string {
	parameterNodeList := functionDeclaratorNode.ChildByFieldName("parameters")
	parameterTypeList := []string{}
	
	if parameterNodeList != nil {
		for i := uint(0); i < parameterNodeList.ChildCount(); i++ {
			parameter := parameterNodeList.Child(i)
			if parameter.Kind() == "parameter_declaration" {
				if parameterTypeNode := parameter.ChildByFieldName("type"); parameterTypeNode != nil {
					rawType := extractText(parameter, source)
					parameterTypeList = append(parameterTypeList, NormalizeType(rawType))
				}
			}
		}
	}

	return parameterTypeList
}

func extractFunctionReturnType(node *ts.Node, source []byte) string {
	returnTypeNode := node.ChildByFieldName("type")
	if returnTypeNode == nil {
		return "void"
	}
	rawType := extractText(returnTypeNode, source)
	return NormalizeType(rawType)
}

func extractIdentifier(node *ts.Node) *ts.Node {
	for node != nil {
		if node.Kind() == "identifier" {
			return node
		}
		node = node.ChildByFieldName("declarator")
	}
	return nil
}

func extractText(node *ts.Node, source []byte) string {
	return string(source[node.StartByte():node.EndByte()])
}
