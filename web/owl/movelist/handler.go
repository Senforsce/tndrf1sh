package movelist

import (
	"fmt"
	"time"

	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
)

var selectDetails = `
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
SELECT ?s ?p ?o WHERE {  ?s rdf:type <http://senforsce.com/o8/brain/SportMovementExercise> .
	?s ?p ?o .
} LIMIT 100`
var panicMessage = "failed sparql query %s"

func ListHandler(c *router.Context) error {
	repo, err1 := sparql.NewRepo("http://localhost:3030/ds",
		sparql.DigestAuth("dba", "dba"),
		sparql.Timeout(time.Millisecond*1500),
	)
	if err1 != nil {
		panic("bad")
	}

	query := selectDetails
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	liste := ListOfSubjects(res.Results.Bindings)
	views := []ViewConfig{}
	for i, _ := range liste {
		view := NewViewConfig(liste[i])
		views = append(views, view)

	}

	return c.Render(HTMX(c, views))
}
