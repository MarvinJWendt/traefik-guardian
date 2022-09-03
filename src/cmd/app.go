package main

import (
	"log"
	"os"
	"time"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/auth"
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html"
)

func main() {
	// Get environment variables
	authDomain := os.Getenv("AUTH_DOMAIN")

	// Setup html templating
	engine := html.New("./html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Setup logger
	app.Use(logger.New())

	app.Use(favicon.New(favicon.Config{
		File: "./html/assets/favicon.ico",
	}))

	// Setup session store
	store := session.New(session.Config{
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:" + auth.SESSION_COOKIE_NAME,
		CookiePath:     "/",
		KeyGenerator:   utils.UUID,
		CookieHTTPOnly: true,
		CookieSameSite: fiber.CookieSameSiteLaxMode,
	})

	// Register routes
	app.Get("/", handlers.IndexRoute(store))
	app.Get("/login", handlers.LoginRoute(store))
	app.Post("/login", handlers.LoginAPI(store))
	app.Get("/logout", handlers.LogoutRoute(store))
	app.Get("/traefik-auth-provider-session-share", handlers.SessionShareRoute())
	app.Get("/check", handlers.CheckRoute(store, authDomain))

	app.Static("/assets", "./html/assets")

	// Start server
	log.Fatal(app.Listen(":80"))
}
