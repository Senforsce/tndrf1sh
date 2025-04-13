package programme

import (
	"fmt"
	"time"

	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
)

func Handler(c *router.Context) error {
	c.Render(HTMX(c))
	return nil
}

var term = "<http://senforsce.com/o8/brain/coach-mj/%s>"
var selectDetails = `
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
SELECT ?s ?p ?o WHERE {  ?s rdf:type <http://senforsce.com/o8/brain/PersonnalizedProgram> .
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

	if err1 != nil {
		panic("bad")
	}
	query := selectDetails
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}
	fmt.Println("///////////////////////////////////////////////")

	fmt.Println(res.Results.Bindings)

	fmt.Println("///////////////////////////////////////////////")

	view := NewViewConfig(res.Results.Bindings)
	views := []ViewConfig{}
	views = append(views, view)
	fmt.Println(views)

	return c.Render(List(views, c))
}
