package model

import "strings"

// TODO: Store the location of the function and raw signature.
type Function struct {
	Name       string   `json:"name"`
	Parameters []string `json:"params"`
	ReturnType string   `json:"return_type"`
}

func (function Function) Signature() string {
	var out strings.Builder
	for _, param := range function.Parameters {
		out.WriteString(param + " -> ")
	}
	out.WriteString(function.ReturnType)
	return out.String()
}

func (function Function) String() string {
	return function.Name + " :: " + function.Signature()
}
