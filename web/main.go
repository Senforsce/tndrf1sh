package main

import (
	"context"
	"embed"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"

	jira "github.com/andygrunwald/go-jira"
	"github.com/joho/godotenv"
	"github.com/senforsce/sparql"
	"github.com/senforsce/telemetry"
	"github.com/senforsce/tndr0cean/router"

	"github.com/senforsce/userconfig"
)

//go:embed static/*
//go:embed static/js/*
//go:embed static/css/*
//go:embed static/img/*
var staticDir embed.FS

func main() {
	tndrf1shContext := context.Background()
	logOptions := telemetry.SenforsceLoggerOptions{
		Context:        tndrf1shContext,
		LogLevel:       "debug",
		LoggerName:     "tndrf1sh",
		ServiceName:    "main",
		ServiceVersion: "0.1.0",
		Verbose:        true,
	}
	tndrf1shLogger, tndrf1shLoggerProvider := telemetry.NewLogger(logOptions)

	defer tndrf1shLoggerProvider.Shutdown(tndrf1shContext)

	app := router.New(tndrf1shLogger)
	Plugs(app)
	userConfig := userconfig.NewUserConfig()
	//	roots := []string{userConfig.StaticRoot, userConfig.SubjectStaticRoot}
	roots := []string{userConfig.StaticRoot, userConfig.SubjectStaticRoot}
	o := router.Options{
		Embedded:    true,
		EmbeddedDir: staticDir,
		StaticRoot:  roots,
	}
	log.Fatal(app.Start(o))

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
	WithAdminSection(app)
	userConfig := userconfig.NewUserConfig()

	app.Plug(WithAuth)
	app.Plug(WithRootDirectories)
	WithSparQlServer(app, userConfig)
	WithHTMXComponents(app)
	WithHTMXRegistry(app)
	WithHTMXPreviews(app)
}

const PAGEPATH = "/"

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
