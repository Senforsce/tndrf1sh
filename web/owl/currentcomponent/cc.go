package currentcomponent

import (
	"fmt"
	"time"

	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

var term = "http://senforsce.com/o8/brain/%s"
var selectDetails = "SELECT ?s ?p ?o WHERE { ?s ?p %s} LIMIT 100"
var panicMessage = "failed sparql query %s"

func Handler(c *router.Context) error {
	repo, err1 := sparql.NewRepo("http://localhost:3030/ds",
		sparql.DigestAuth("dba", "dba"),
		sparql.Timeout(time.Millisecond*1500),
	)
	if err1 != nil {
		panic("bad")
	}
	query := fmt.Sprintf(selectDetails, "?o")
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	so := meta.ShouldInspect(c.Request.URL)

	return c.Render(HTMX(res.Results.Bindings, meta.Config{
		TargetSelector: "",
		Inspect:        so,
	}, c))
}
