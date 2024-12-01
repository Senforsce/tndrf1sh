package abdoulsylogo

import (
	"github.com/senforsce/tndr0cean/router"
)

func Preview(c *router.Context) error {
	return Handler(c)
}
