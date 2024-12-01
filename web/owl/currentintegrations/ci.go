package currentintegrations

import (
	"fmt"

	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

var term = "http://senforsce.com/o8/brain/%s"
var selectDetails = "SELECT ?s ?p ?o WHERE { ?s a %s}"
var panicMessage = "failed sparql query %s"

func Handler(c *router.Context) error {
	sparql := c.Get("sparql").(sparql.Repo)
	o8Type := fmt.Sprintf(term, "sen:Thunderf1shIntegration")
	query := fmt.Sprintf(selectDetails, o8Type)
	res, err := sparql.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	so := meta.ShouldInspect(c.Request.URL)

	return c.Render(HTMX(res.Results.Bindings, meta.Config{
		TargetSelector: "",
		Inspect:        so,
	}, c))
}
