package main

import (
	"github.com/MarvinJWendt/simple-forward-auth/src/internal/pkg/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"

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

	// Setup session store
	store := session.New(session.Config{
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:simple_forward_auth_session_id",
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
	app.Get("/simple-forward-auth-session-share", handlers.SessionShareRoute())
	app.Get("/check", handlers.CheckRoute(store, authDomain))

	// Start server
	log.Fatal(app.Listen(":80"))
}
