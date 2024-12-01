package main

import (
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/owl/abdoulsylogo" 
)

func WithHXMXComponents(app *router.Tndr0cean) error {
	app.Get("/preview/abdoulsylogo", abdoulsylogo.Preview) 
	

	return nil
}
	