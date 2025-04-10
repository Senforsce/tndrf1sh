package moves

import (
	"strings"

	"github.com/senforsce/sparql"
)

type ViewConfig struct {
	Name string
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

func NewViewConfig(results []map[string]sparql.Binding) *ViewConfig {
	return &ViewConfig{
		Name: FindObjectValueByPredicate("MoveExerciseName", results)["o"].Value,
	}
}
