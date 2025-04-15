package movelist

import (
	"github.com/senforsce/tndr0cean/router"
)

func Preview(c *router.Context) error {
	return c.Render(HTMX(c, []ViewConfig{}))
}
