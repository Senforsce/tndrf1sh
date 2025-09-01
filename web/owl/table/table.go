package table

import (
	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
)

func Handler(c *router.Context, res []map[string]sparql.Binding) error {
	c.Render(HTMX(c, res))
	return nil
}
