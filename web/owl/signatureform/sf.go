package signatureform

import (
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

func Handler(c *router.Context) error {

	so := meta.ShouldInspect(c.Request.URL)

	return c.Render(HTMX(meta.Config{
		TargetSelector: "",
		Inspect:        so,
	}, c))
}
