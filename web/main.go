package main

import (
	"fmt"
	"log"
	"os"
	"time"

	jira "github.com/andygrunwald/go-jira"
	"github.com/joho/godotenv"
	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout"
)

func main() {
	app := router.New()
	Plugs(app)
	log.Fatal(app.Start())

}

func init() {
	if envErr := godotenv.Load(".env"); envErr != nil {
		fmt.Println(".env file missing")
	}
}

func Plugs(app *router.Tndr0cean) {
	WithO8(app)

	app.Plug(WithAuth)
	WithSparQlServer(app)
	WithHTMXComponents(app)
	WithHTMXRegistry(app)
	WithHTMXServer(app)
	// WithHTMXPreviews(app)
}

func WithHTMXServer(app *router.Tndr0cean) func(h router.Handler) {
	app.Get("/", layout.Handler)
	// Other routes will be injected with tree-shaking on build inside ./injected_routes.go
	return nil
}

func WithSparQlServer(app *router.Tndr0cean) func(h router.Handler) {
	repo, err := sparql.NewRepo("http://localhost:3030/ds",
		sparql.DigestAuth("dba", "dba"),
		sparql.Timeout(time.Millisecond*1500),
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
	return func(c *router.Context) error {
		c.Set("username", "Senforsce")
		c.Set("password", "test")

		return h(c)
	}
}

func jiraClient() *jira.Client {
	jt := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USER"),
		Password: os.Getenv("JIRA_TOKEN"),
	}

	client, err := jira.NewClient(jt.Client(), os.Getenv("JIRA_URL"))
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
