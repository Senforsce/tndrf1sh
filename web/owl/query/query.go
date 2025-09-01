package query

import "github.com/senforsce/tndr0cean/router"

func Handler(c *router.Context, qr string) error {
	c.Render(HTMX(c, qr))
	return nil
}
