package object

import (
	rds "github.com/senforsce/auth/exp/store/redis"
	"github.com/senforsce/auth/pkg/auth"
	"github.com/senforsce/fh/forms/tripleforms"
	"github.com/senforsce/htmx/hx"
	"github.com/senforsce/owl/logged"
	"github.com/senforsce/tndr0cean/router"
)

func Routes(app *router.Tndr0cean, sessionStore *rds.RedisSessionManager, authClient *auth.Client) {
	hx.Route(app, HXEDITOBJECTPATH, logged.Logged(sessionStore, authClient, EditHandler))
	hx.Route(app, HXVIEWOBJECTPATH, logged.Logged(sessionStore, authClient, ViewHandler))
	hx.Route(app, HXDELETEOBJECTPATH, logged.Logged(sessionStore, authClient, tripleforms.DeleteTripleHandler))

	hx.Route(app, tripleforms.HXDELETEPATH, logged.Logged(sessionStore, authClient, tripleforms.ActuallyDeleteTriple))
	hx.Route(app, tripleforms.HXGETPATH, logged.Logged(sessionStore, authClient, tripleforms.EditTripleHandler))
	hx.RoutePost(app, tripleforms.HXPOSTPATH, logged.Logged(sessionStore, authClient, tripleforms.HandleObjectValue))
}
