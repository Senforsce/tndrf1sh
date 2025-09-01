package user

import "github.com/senforsce/tndr0cean/router"

func Handler(c *router.Context) error {
	c.Render(HTMX(c))

	return nil
}

func EditHandler(c *router.Context) error {
	c.Render(HTMX(c))

	return nil
}

func DeleteHandler(c *router.Context) error {
	c.Render(HTMX(c))

	return nil
}
