package predicate

import (
	rds "github.com/senforsce/auth/exp/store/redis"
	"github.com/senforsce/auth/pkg/auth"
	"github.com/senforsce/fh/forms/tripleforms"
	"github.com/senforsce/htmx/hx"
	"github.com/senforsce/owl/logged"
	"github.com/senforsce/tndr0cean/router"
)

func Routes(app *router.Tndr0cean, sessionStore *rds.RedisSessionManager, authClient *auth.Client) {
	hx.Route(app, HXEDITPATH, logged.Logged(sessionStore, authClient, EditHandler))
	hx.Route(app, HXVIEWPATH, logged.Logged(sessionStore, authClient, ViewHandler))
	hx.Route(app, HXADDPREDICATEPATH, logged.Logged(sessionStore, authClient, AddformViewHandler))
	hx.Route(app, tripleforms.HXADDGETPATH, logged.Logged(sessionStore, authClient, tripleforms.AddTripleHandler))
	hx.RoutePost(app, tripleforms.HXADDPOSTPATH, logged.Logged(sessionStore, authClient, tripleforms.HandlePredicateExerciceValue))

}
