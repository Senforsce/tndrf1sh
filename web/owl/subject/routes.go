package subject

import (
	rds "github.com/senforsce/auth/exp/store/redis"
	"github.com/senforsce/auth/pkg/auth"
	"github.com/senforsce/htmx/hx"
	"github.com/senforsce/owl/logged"
	"github.com/senforsce/tndr0cean/router"
)

func Routes(app *router.Tndr0cean, sessionStore *rds.RedisSessionManager, authClient *auth.Client) {
	hx.Route(app, HXEDITPATH, logged.Logged(sessionStore, authClient, EditHandler))
	hx.Route(app, HXVIEWPATH, logged.Logged(sessionStore, authClient, ViewHandler))

}
