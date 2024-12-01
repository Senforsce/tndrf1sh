package layout

import (
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

func Handler(c *router.Context) error {
	conf := meta.NewComponentConfig(c)
	return c.Render(Base(conf.Ctx))
}
