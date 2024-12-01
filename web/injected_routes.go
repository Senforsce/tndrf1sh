package main

import (
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/owl/componentpreview"
	"github.com/senforsce/tndrf1sh/web/owl/currentcomponent"
	dirls "github.com/senforsce/tndrf1sh/web/owl/dirLs"
	"github.com/senforsce/tndrf1sh/web/owl/texteditor"
)

func WithHTMXComponents(app *router.Tndr0cean) error {
	app.Get("/owl/*filepath", dirls.Handler)
	app.Get("/file/open/*filepath", texteditor.Handler)
	app.Post("/owl/current/*filepath", componentpreview.Handler)
	app.Post("/components", currentcomponent.Handler)

	return nil
}
