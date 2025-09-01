package programmetypedrag

import (
	"strings"

	"github.com/senforsce/sparql"
)

type ViewConfig struct {
	ProgramTypeName string
	Comment         string
	Subject         string
	Illustration    string
	BgColor         string
	ShortDefinition string
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

func ListOfSubjects(results []map[string]sparql.Binding) map[string][]map[string]sparql.Binding {
	var toReturn = make(map[string][]map[string]sparql.Binding)

	for _, term := range results {
		// Check if the predicate (p) matches the needle
		if _, ok := term["s"]; ok {
			val := term["s"].Value

			if len(toReturn[val]) < 1 {
				toReturn[val] = []map[string]sparql.Binding{}
			}
			toReturn[val] = append(toReturn[val], term)
		}
	}

	return toReturn
}

func NewViewConfig(results []map[string]sparql.Binding) ViewConfig {
	return ViewConfig{
		ProgramTypeName: FindObjectValueByPredicate("ProgramTypeName", results)["o"].Value,
		Comment:         FindObjectValueByPredicate("comment", results)["o"].Value,
		Subject:         FindObjectValueByPredicate("ProgramTypeName", results)["s"].Value,
		Illustration:    FindObjectValueByPredicate("ProgramTypePictureIllustration", results)["o"].Value,
		BgColor:         FindObjectValueByPredicate("ProgramTypeColor", results)["o"].Value,
		ShortDefinition: FindObjectValueByPredicate("ProgramTypeShortDefinition", results)["o"].Value,
	}
}
