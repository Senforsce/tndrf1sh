package main

import (
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/owl/htm/about"
	"github.com/senforsce/owl/htm/ocean"
	"github.com/senforsce/tndrf1sh/web/owl/abdoulsylogo"
	
)

func WithHTMXPreviews(app *router.Tndr0cean) error {
	app.Get("/preview/about", about.Preview) 
	app.Get("/preview/ocean", ocean.Preview) 
	app.Get("/preview/abdoulsylogo", abdoulsylogo.Preview) 
	

	return nil
}
	