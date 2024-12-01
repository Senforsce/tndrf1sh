package abdoulsylogo

import (
	"github.com/senforsce/tndr0cean/router"
)

func Handler(c *router.Context) error {
	return c.Render(HTMX(c))
}
