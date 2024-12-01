package cliquablelist

import (
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

func Handler(c *router.Context) error {
	Items := c.Get("components")
	targetSelector := c.Get("components_target")

	so := meta.ShouldInspect(c.Request.URL)

	return c.Render(HTMX(Items.([]Item), meta.Config{
		TargetSelector: targetSelector.(string),
		Inspect:        so,
	}, c))
}
