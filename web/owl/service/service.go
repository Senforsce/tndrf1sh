package service

import "github.com/senforsce/tndr0cean/router"

func Handler(c *router.Context) error {
	c.Render(HTMX(c))
	return nil

}
