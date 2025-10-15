package main

import (
	"fmt"

	"github.com/senforsce/auth/exp/config"

	redis "github.com/redis/go-redis/v9"
	authhandler "github.com/senforsce/auth/exp/handler/auth"
	rds "github.com/senforsce/auth/exp/store/redis"
	"github.com/senforsce/auth/pkg/auth"
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout"
)

func WithAuthenticationSystem(app *router.Tndr0cean, redisClient *redis.Client, sessionStore *rds.RedisSessionManager, authClient *auth.Client, cfg *config.Config) {

	serverAddr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	authStore := rds.NewAuthRedisManager(redisClient)

	authHandler := authhandler.New(cfg,
		serverAddr,
		authClient,
		authStore,
		sessionStore,
	)

	app.Get("/login", authHandler.RenderLoginPage)
	app.Get("/mj/logout", Logout(sessionStore, authClient, layout.Handler))
	app.Get("/auth/login", authHandler.RedirectToKeycloak)
	app.Get("/login-keycloak", authHandler.RedirectToKeycloak)
	app.Get("/auth/callback", authHandler.Callback)
}
