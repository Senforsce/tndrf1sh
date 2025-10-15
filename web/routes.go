package main

import (
	rds "github.com/senforsce/auth/exp/store/redis"
	"github.com/senforsce/auth/pkg/auth"
	"github.com/senforsce/tndr0cean/router"
)

func Routes(app *router.Tndr0cean, sessionStore *rds.RedisSessionManager, authClient *auth.Client) {
	//HTMX FORMS FOR USER
	//app.Get(PAGEPATH, logged.Logged(sessionStore, authClient, layout.Handler))
}
