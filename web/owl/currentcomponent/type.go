package currentcomponent

import (
	"strings"

	"github.com/senforsce/sparql"
)

type ViewConfig struct {
	Title            string
	FrTitle          string
	FiletypeSymbol   string //TODO point to fontawesome or SVG icons
	Description      string
	FrDescription    string
	FilePathFromRoot string
	Project          string
	Package          string
}

//naive find triple

func FindObjectValueByPredicate(needle string, haystack []map[string]sparql.Binding) map[string]sparql.Binding {
	for _, term := range haystack {
		// Check if the predicate (p) matches the needle
		if val, exists := term["p"]; exists {
			if strings.Contains(val.Value, needle) {
				// If a match is found, return the object (o) and true
				return term
			}
		}
	}
	// If no match is found, return an empty map and false
	return map[string]sparql.Binding{}
}

func getSymbol(extension string) string {
	switch extension {
	case "t1":
		return "✍"
	case "_t1go":
		return "🦾"
	case "go":
		return "🥑"
	default:
		return "🌶"
	}
}

func NewViewConfig(results []map[string]sparql.Binding) *ViewConfig {
	return &ViewConfig{
		Project:          "coach-mj",
		Title:            FindObjectValueByPredicate("t1packageName", results)["o"].Value,
		FrTitle:          FindObjectValueByPredicate("t1packageName", results)["o"].Value,
		Package:          FindObjectValueByPredicate("t1packageName", results)["o"].Value,
		FiletypeSymbol:   getSymbol(FindObjectValueByPredicate("t1FileExtension", results)["o"].Value),
		FilePathFromRoot: FindObjectValueByPredicate("filepathFromRoot", results)["o"].Value,
		Description:      FindObjectValueByPredicate("comment", results)["o"].Value,
		FrDescription:    FindObjectValueByPredicate("comment", results)["o"].Value,
	}
}
