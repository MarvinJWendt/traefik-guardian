package main

import (
	"log"
	"os"
	"time"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/auth"
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/config"
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html"
	"github.com/pterm/pterm"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func main() {
	pterm.DefaultBox.Println(
		"                Traefik Auth Provider" +
			"\n" +
			"https://github.com/MarvinJWendt/traefik-auth-provider",
	)

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})

	logrus.Debug("initializing config ", config.CONFIG_DIR+"/"+config.CONFIG_FILE)
	err := config.Initzialize()
	if err != nil {
		logrus.Fatal("Failed to initialize config", zap.Error(err))
	}

	// Get environment variables
	authDomain := os.Getenv("AUTH_DOMAIN")

	// Setup html templating
	logrus.Debug("initializing html templating")
	engine := html.New("./html", ".html")

	logrus.Debug("creating web server instance")
	app := fiber.New(fiber.Config{
		Views:                 engine,
		DisableStartupMessage: true,
	})

	logrus.Debug("adding favicon middleware")
	app.Use(favicon.New(favicon.Config{
		File: "./html/assets/favicon.ico",
	}))

	// Setup session store
	logrus.Debug("initializing session store")
	store := session.New(session.Config{
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:" + auth.SESSION_COOKIE_NAME,
		CookiePath:     "/",
		KeyGenerator:   utils.UUID,
		CookieHTTPOnly: true,
		CookieSameSite: fiber.CookieSameSiteLaxMode,
	})

	// Register routes
	logrus.Debug("registering routes")
	app.Get("/", handlers.IndexRoute(store))
	app.Get("/login", handlers.LoginRoute(store))
	app.Post("/login", handlers.LoginAPI(store))
	app.Get("/logout", handlers.LogoutRoute(store))
	app.Get("/traefik-auth-provider-session-share", handlers.SessionShareRoute())
	app.Get("/check", handlers.CheckRoute(store, authDomain))

	logrus.Debug("registering static file server for assets")
	app.Static("/assets", "./html/assets")

	// Start server
	logrus.Debug("starting web server")
	log.Fatal(app.Listen(":80"))
}
