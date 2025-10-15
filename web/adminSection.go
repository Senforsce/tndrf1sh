package main

import (
	"context"
	_ "embed"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"

	redis "github.com/redis/go-redis/v9"

	"github.com/senforsce/auth/exp/config"
	"github.com/senforsce/auth/exp/middleware"
	rds "github.com/senforsce/auth/exp/store/redis"
	"github.com/senforsce/auth/pkg/auth"

	"github.com/senforsce/coachmj/web/pages"

	"github.com/senforsce/tndr0cean/router"
)

// File [./web/admin/partieadmin.go](https://github.com/senforsce/coachmj/web/admin/):
//
// Fonction WithAdminSection :
//
// # Ajoute les capabilités pour le site web administrateur français utilisé pour coach-mj
//
// ce module ajoute les routes ["/mj/dashboard", "/mj/userinfo", "/calendar", "/pr", "/preferred"]
func WithAdminSection(app *router.Tndr0cean) {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatalf("failed to load and parse config : %v", err)
		return
	}
	ctx := context.Background()

	authOptions := []auth.Option{
		auth.WithClientSecret(cfg.Auth.ClientSecret),
		auth.WithRealmKeycloak(cfg.Auth.Realm),
	}
	authClient, err := auth.New(
		ctx,
		cfg.Auth.BaseURL,
		cfg.Auth.ClientID,
		cfg.Auth.RedirectURL,
		authOptions...,
	)
	if err != nil {
		log.Fatalf("Failed to initialize auth client : %v", err)
		return
	}
	//Redis
	redisClient := redis.NewClient(&cfg.RedisConfig)
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
		return
	}

	sessionStore := rds.NewSessionRedisManager(redisClient)

	//renderHandler := render.New(cfg)
	WithAuthenticationSystem(app, redisClient, sessionStore, authClient, cfg)
	WithDynamicHTMXComponents(app, sessionStore, authClient)
	//WithDynamicFormHandlers(app, sessionStore, authClient)
	// carnetdebord.Routes(app, sessionStore, authClient)
	// programme.Routes(app, sessionStore, authClient)
	// movelist.Routes(app, sessionStore, authClient)
	// calendrier.Routes(app, sessionStore, authClient)
	// recordsperso.Routes(app, sessionStore, authClient)
	// mouvementbibliotheque.Routes(app, sessionStore, authClient)
	// uploads.Routes(app, sessionStore, authClient)
	// mouvementinfos.Routes(app, sessionStore, authClient)
	// mesinfos.Routes(app, sessionStore, authClient)
	// pr.Routes(app, sessionStore, authClient)

	//HTMX FORMS FOR USER PROFILE

}

func Logout(sessionStore *rds.RedisSessionManager, authClient *auth.Client, handler func(*pages.Ctx) error) router.Handler {
	return router.WithMid(middleware.AuthMiddleware(sessionStore, authClient), handler)
}
