package movement

import (
	"fmt"
	"strings"
	"time"

	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
)

var term = "<%s>"
var selectDetails = "SELECT ?s ?p ?o WHERE {  %s ?p ?o . } LIMIT 100"
var panicMessage = "failed sparql query %s"

func Handler(c *router.Context) error {
	repo, err1 := sparql.NewRepo("http://localhost:3030/ds",
		sparql.DigestAuth("dba", "dba"),
		sparql.Timeout(time.Millisecond*1500),
	)
	if err1 != nil {
		panic("bad")
	}

	fmt.Println("a", c.Query("moveId"))
	fmt.Println("b", c.Param("moveId"))
	fmt.Println("c", c.FormValue("moveId"))
	fmt.Println("d", c.Request.PathValue("moveId"))

	moveId := strings.Replace(c.Param("moveId"), "/", "", 1)

	query := fmt.Sprintf(selectDetails, fmt.Sprintf(term, moveId))
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf("bad", query))
	}
	fmt.Println(query)

	fmt.Println(res.Results.Bindings)
	em := NewExerciseMove(res.Results.Bindings)
	c.Render(HTMX(c, em))
	return nil

}

type ExerciseMove struct {
	Name        string
	Description string
	GifPath     string //TODO point to fontawesome or SVG icons
	Tips        string
	Groupes     string
	Credit      string
}

func NewExerciseMove(results []map[string]sparql.Binding) *ExerciseMove {
	return &ExerciseMove{
		Name:        FindObjectValueByPredicate("MoveExerciceName", results)["o"].Value,
		Description: FindObjectValueByPredicate("MoveDescription", results)["o"].Value,
		GifPath:     FindObjectValueByPredicate("MJMoveGifMediaDescription", results)["o"].Value,
		Tips:        FindObjectValueByPredicate("MoveTipsFlattened", results)["o"].Value,
		Groupes:     FindObjectValueByPredicate("WorksMuscularGroupsFlattened ", results)["o"].Value,
		Credit:      FindObjectValueByPredicate("MoveGifDescriptionCredits", results)["o"].Value,
	}
}

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
