package userlistdrag

import (
	"strings"

	"github.com/senforsce/sparql"
)

type ViewConfig struct {
	Firstname         string
	Lastname          string
	ContactEmail      string
	Telephone         string
	ProfilePicture    string
	Subject           string
	HasProfilePicture ProfilePicture
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

func NewProfilePicture(res []map[string]sparql.Binding, subject string) ProfilePicture {
	return ProfilePicture{
		Filename: sparql.GetValue("ProfilePictureFilename", res),
		Filepath: sparql.GetValue("ProfilePictureFilepath", res),
		Salt:     sparql.GetValue("ProfilePictureSalt", res),
		Subject:  subject,
	}
}

type ProfilePicture struct {
	Filename string
	Data     string
	Filepath string
	Salt     string
	Subject  string
}

type CompleteMainInfo struct {
	Birthdaydate      string
	BelongsTo         string
	Subject           string
	Object            string
	PrChartDataLabels string
	PrChartData       string
	Firstname         string
	Lastname          string
	Email             string
	ProfilePicture    string
	HasProfilePicture ProfilePicture
	Telephone         string
}

func NewCompleteMainInfo(results []map[string]sparql.Binding) CompleteMainInfo {
	return CompleteMainInfo{
		BelongsTo:         sparql.GetValue("main", results),
		Subject:           sparql.GetValue("s", results),
		Birthdaydate:      sparql.GetValue("Birthdaydate", results),
		Firstname:         sparql.GetValue("Firstname", results),
		Lastname:          sparql.GetValue("Lastname", results),
		Email:             sparql.GetValue("Email", results),
		HasProfilePicture: NewProfilePicture(results, sparql.GetValue("s", results)),
		Telephone:         sparql.GetValue("Telephone", results),
	}
}

func NewViewConfig(results []map[string]sparql.Binding) ViewConfig {
	return ViewConfig{
		Firstname:         FindObjectValueByPredicate("hasFirstname", results)["o"].Value,
		Lastname:          FindObjectValueByPredicate("hasLastname", results)["o"].Value,
		ContactEmail:      FindObjectValueByPredicate("hasContactEmail", results)["o"].Value,
		Telephone:         FindObjectValueByPredicate("hasTelephone", results)["o"].Value,
		ProfilePicture:    FindObjectValueByPredicate("hasProfilePicture", results)["o"].Value,
		Subject:           FindObjectValueByPredicate("hasFirstname", results)["s"].Value,
		HasProfilePicture: NewProfilePicture(results, FindObjectValueByPredicate("hasFirstname", results)["s"].Value),
	}
}
