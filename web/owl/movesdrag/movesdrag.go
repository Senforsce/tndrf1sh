package movesdrag

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
)

var term = "<http://senforsce.com/o8/brain/coach-mj/%s>"
var selectDetails = `
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
SELECT ?s ?p ?o WHERE {  ?s rdf:type <http://senforsce.com/o8/brain/SportMovement> .
	?s ?p ?o .
} LIMIT 1000`
var panicMessage = "failed sparql query %s"

func Handler(c *router.Context) error {
	repo, err1 := sparql.NewRepo("http://localhost:3030/ds",
		sparql.DigestAuth("dba", "dba"),
		sparql.Timeout(time.Millisecond*1500),
	)
	if err1 != nil {
		panic("bad")
	}

	skip := strings.Replace(c.Param("skip"), "/", "", 1)

	i, err := strconv.Atoi(skip)
	if err != nil {
		panic("bad")
	}
	skipValue := i

	query := selectDetails
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	liste := ListOfSubjects(res.Results.Bindings)
	views := []ViewConfig{}
	for i, _ := range liste {
		id, _ := gonanoid.New()

		view := NewViewConfig(liste[i], id)
		views = append(views, view)

	}
	return c.Render(HTMX(views, c, skipValue))
}
