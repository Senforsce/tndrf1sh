package main

import (
	"fmt"
	"log"
	"os"
	"time"

	jira "github.com/andygrunwald/go-jira"
	"github.com/joho/godotenv"
	"github.com/senforsce/fh/handlers"
	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout"
	"github.com/senforsce/tndrf1sh/web/owl/abonnement"
	"github.com/senforsce/tndrf1sh/web/owl/blogtexteditor"
	"github.com/senforsce/tndrf1sh/web/owl/programme"

	"github.com/senforsce/tndrf1sh/web/owl/movement"
	"github.com/senforsce/tndrf1sh/web/owl/moves"
	"github.com/senforsce/tndrf1sh/web/owl/service"
	"github.com/senforsce/tndrf1sh/web/owl/user"
	"github.com/senforsce/tndrf1sh/web/owl/userlist"

	"github.com/senforsce/userconfig"
)

func main() {
	app := router.New()
	Plugs(app)
	userConfig := userconfig.NewUserConfig()

	log.Fatal(app.Start(userConfig.StaticRoot))

}

func init() {
	if envErr := godotenv.Load(".env"); envErr != nil {
		fmt.Println(".env file missing")
	}
}

func Plugs(app *router.Tndr0cean) {
	WithO8(app)
	WithO8MJ(app)
	WithBootstrap(app)
	userConfig := userconfig.NewUserConfig()

	app.Plug(WithAuth)
	app.Plug(WithRootDirectories)
	WithSparQlServer(app, userConfig)
	WithHTMXComponents(app)
	WithHTMXRegistry(app)
	WithHTMXServer(app)
	WithHTMXPreviews(app)
}

func WithHTMXServer(app *router.Tndr0cean) func(h router.Handler) {
	app.Get("/", layout.Handler)
	app.Get("/forms/newMovement", handlers.FormNewMovement)
	app.Get("/blog/new", blogtexteditor.Handler)
	app.Get("/programme/new", programme.Handler)

	app.Post("/process/newMove", handlers.HandleCreateNewMovement)
	app.Get("/forms/newUser", handlers.FormNewUser)
	app.Post("/process/newUser", handlers.HandleCreateNewUser)
	app.Get("/mj/moves", moves.Handler)
	app.Get("/mj/userlist", userlist.Handler)
	app.Get("/mj/user/*userName", user.Handler)
	app.Get("/mj/movement/*moveId", movement.Handler)
	app.Get("/mj/abonnement/*abonnementName", abonnement.Handler)
	app.Get("/mj/service/*serviceName", service.Handler)

	// Other routes will be injected with tree-shaking on build inside ./injected_routes.go
	return nil
}

func WithSparQlServer(app *router.Tndr0cean, userConfig *userconfig.UserConfig) func(h router.Handler) {
	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	r, err := repo.Query("SELECT * WHERE { ?s ?p ?o } LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)
	// app.Post(app.G("path:SparQlCurrentIntegrations"), currentintegrations.Handler)
	// app.Post(app.G("path:SparQlCurrentComponent"), currentcomponent.Handler)

	return nil

}

func WithAuth(h router.Handler) router.Handler {
	userConfig := userconfig.NewUserConfig()

	return func(c *router.Context) error {
		c.Set(userConfig.AuthUsernameKey, userConfig.AuthPasswordValue)
		c.Set(userConfig.AuthPasswordKey, userConfig.AuthPasswordValue)

		return h(c)
	}
}

func WithRootDirectories(h router.Handler) router.Handler {
	userConfig := userconfig.NewUserConfig()

	return func(c *router.Context) error {
		c.Set("O8ROOT", userConfig.O8Root)
		c.Set("CSSROOT", userConfig.CSSRoot)

		return h(c)
	}
}

func jiraClient(config userconfig.UserConfig) *jira.Client {
	jt := jira.BasicAuthTransport{
		Username: os.Getenv(config.TicketSystemUsername),
		Password: os.Getenv(config.TicketSystemToken),
	}

	client, err := jira.NewClient(jt.Client(), os.Getenv(config.TicketSystemUrl))
	if err != nil {
		fmt.Println(err)
	}

	me, _, err := client.User.GetSelf()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(me)

	return client
}
