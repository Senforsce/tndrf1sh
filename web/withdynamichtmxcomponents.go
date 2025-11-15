package main

import (
	rds "github.com/senforsce/auth/exp/store/redis"
	"github.com/senforsce/auth/pkg/auth"
	"github.com/senforsce/fh/handlers"
	"github.com/senforsce/htmx/hx"
	"github.com/senforsce/owl/connected/movelist"
	"github.com/senforsce/owl/logged"
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout"
	"github.com/senforsce/tndrf1sh/web/owl/abonnement"
	"github.com/senforsce/tndrf1sh/web/owl/blogtexteditor"
	"github.com/senforsce/tndrf1sh/web/owl/bookingcalendar"
	"github.com/senforsce/tndrf1sh/web/owl/movement"
	"github.com/senforsce/tndrf1sh/web/owl/moves"
	"github.com/senforsce/tndrf1sh/web/owl/movesdrag"
	"github.com/senforsce/tndrf1sh/web/owl/object"
	"github.com/senforsce/tndrf1sh/web/owl/predicate"
	"github.com/senforsce/tndrf1sh/web/owl/programme"
	"github.com/senforsce/tndrf1sh/web/owl/programmetypedrag"
	"github.com/senforsce/tndrf1sh/web/owl/queryandtable"
	"github.com/senforsce/tndrf1sh/web/owl/service"
	"github.com/senforsce/tndrf1sh/web/owl/subject"
	"github.com/senforsce/tndrf1sh/web/owl/user"
	"github.com/senforsce/tndrf1sh/web/owl/userlist"
	"github.com/senforsce/tndrf1sh/web/owl/userlistdrag"
)

func WithDynamicHTMXComponents(app *router.Tndr0cean, sessionStore *rds.RedisSessionManager, authClient *auth.Client) {
	hx.Route(app, "/mj/dashboard", logged.Logged(sessionStore, authClient, layout.Handler))
	hx.Route(app, "/", logged.Logged(sessionStore, authClient, layout.Handler))
	hx.Route(app, "/forms/newMovement", logged.Logged(sessionStore, authClient, handlers.FormNewMovement))
	hx.Route(app, "/blog/new", logged.Logged(sessionStore, authClient, blogtexteditor.Handler))
	hx.Route(app, "/calendar/view", logged.Logged(sessionStore, authClient, bookingcalendar.Handler))
	hx.Route(app, "/calendar/table", logged.Logged(sessionStore, authClient, bookingcalendar.TableHandler))

	hx.Route(app, "/programme/new", logged.Logged(sessionStore, authClient, programme.Handler))
	hx.Route(app, "/programme/table", logged.Logged(sessionStore, authClient, programme.TableHandler))
	hx.Route(app, "/forms/newUser", logged.Logged(sessionStore, authClient, handlers.FormNewUser))
	hx.Route(app, "/mj/moves", logged.Logged(sessionStore, authClient, moves.Handler))
	hx.Route(app, "/mj/movelist", logged.Logged(sessionStore, authClient, movelist.ListHandler))
	hx.Route(app, "/mj/programmes", logged.Logged(sessionStore, authClient, programme.ListHandler))
	hx.Route(app, "/mj/programmetable", logged.Logged(sessionStore, authClient, programme.ListHandler))
	hx.Route(app, "/mj/queryandtable", logged.Logged(sessionStore, authClient, queryandtable.Handler))
	hx.Route(app, "/mj/userlist", logged.Logged(sessionStore, authClient, userlist.Handler))
	hx.Route(app, "/mj/user/*userIRI", logged.Logged(sessionStore, authClient, user.Handler))
	hx.Route(app, "/mj/movement/show/*moveId", logged.Logged(sessionStore, authClient, movement.Handler))
	hx.Route(app, "/mj/movement/edit/*moveId", logged.Logged(sessionStore, authClient, movement.EditHandler))
	hx.Route(app, "/mj/movement/delete/*moveId", logged.Logged(sessionStore, authClient, movement.DeleteHandler))

	hx.Route(app, "/mj/abonnement/*abonnementName", logged.Logged(sessionStore, authClient, abonnement.Handler))
	hx.Route(app, "/mj/service/*serviceName", logged.Logged(sessionStore, authClient, service.Handler))
	hx.Route(app, "/mj/movesdrag/*skip", logged.Logged(sessionStore, authClient, movesdrag.Handler))
	hx.Route(app, "/mj/userlistdrag", logged.Logged(sessionStore, authClient, userlistdrag.Handler))
	hx.Route(app, "/mj/programmetypedrag", logged.Logged(sessionStore, authClient, programmetypedrag.Handler))

	app.Post("/programme/process", logged.Logged(sessionStore, authClient, handlers.HandleCreateNewProgramme))
	app.Post("/process/newMove", logged.Logged(sessionStore, authClient, handlers.HandleCreateNewMovement))
	app.Post("/process/newUser", logged.Logged(sessionStore, authClient, handlers.HandleCreateNewUser))
	app.Put("/mj/programme/*programmeName", logged.Logged(sessionStore, authClient, programme.ListHandler))
	app.Delete("/mj/programme/*programmeName", logged.Logged(sessionStore, authClient, programme.ListHandler))
	app.Put("/mj/user/*userIRI", logged.Logged(sessionStore, authClient, user.EditHandler))
	app.Delete("/mj/user/*userIRI", logged.Logged(sessionStore, authClient, user.DeleteHandler))
	app.Post("/process/editMove", logged.Logged(sessionStore, authClient, handlers.HandleEditMovement))

	programme.Routes(app, sessionStore, authClient)
	subject.Routes(app, sessionStore, authClient)
	predicate.Routes(app, sessionStore, authClient)
	object.Routes(app, sessionStore, authClient)

	// Other routes will be injected with tree-shaking on build inside ./injected_routes.go
}
