package search

import (
	"catcatgo/internal/model"
	"strings"
)

// TODO: Add support for regex.
// TODO: Add support for parameter placeholders.
// TODO: Fuzzy search.
func Linear(sampleSpace []model.Function, query string) []model.Function {
	resultSet := []model.Function{}
	for _, function := range sampleSpace {
		if strings.Contains(function.Name, query) ||
			strings.Contains(function.Signature(), query) {
			resultSet = append(resultSet, function)
		}
	}
	return resultSet
}
