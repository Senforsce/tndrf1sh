package movesdrag

import (
	"strings"

	"github.com/senforsce/sparql"
)

type ViewConfig struct {
	MoveDescription              string
	WorksMuscularGroupsFlattened string
	MoveGifDescriptionCredits    string
	MJMoveGifMediaDescription    string
	MoveTipsFlattened            string
	MoveExerciceName             string
	Subject                      string
	MovePoster                   string
	MoveIcon                     string
	NewId                        string
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

func NewViewConfig(results []map[string]sparql.Binding, tempId string) ViewConfig {
	return ViewConfig{
		MoveDescription:              FindObjectValueByPredicate("MoveDescription", results)["o"].Value,
		MoveGifDescriptionCredits:    FindObjectValueByPredicate("MoveGifDescriptionCredits", results)["o"].Value,
		WorksMuscularGroupsFlattened: FindObjectValueByPredicate("WorksMuscularGroupsFlattened", results)["o"].Value,
		MJMoveGifMediaDescription:    FindObjectValueByPredicate("MJMoveGifMediaDescription", results)["o"].Value,
		MoveTipsFlattened:            FindObjectValueByPredicate("MoveTipsFlattened", results)["o"].Value,
		MoveExerciceName:             FindObjectValueByPredicate("MoveExerciceName", results)["o"].Value,
		MovePoster:                   FindObjectValueByPredicate("MovePosterPicture", results)["o"].Value,
		MoveIcon:                     FindObjectValueByPredicate("MoveIcon", results)["o"].Value,
		NewId:                        tempId,
		Subject:                      FindObjectValueByPredicate("MoveExerciceName", results)["s"].Value,
	}
}
