package main

import (
	"os"
	"time"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/middleware/fiberlog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/sirupsen/logrus"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/api"
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/auth"
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/config"
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/db"
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/handlers"
)

func main() {
	pterm.DefaultBox.Println(
		putils.CenterText(
			"Traefik Auth Provider\n" +
				"https://github.com/MarvinJWendt/traefik-auth-provider",
		),
	)

	time.Sleep(time.Millisecond) // Don't ask why, but this fixes the docker-compose log

	logrus.SetFormatter(&logrus.TextFormatter{})

	err := config.Initialize()
	if err != nil {
		logrus.Fatal("Failed to initialize config: ", err)
	}

	if config.Debug.ToBool() {
		logrus.SetLevel(logrus.DebugLevel)
	}

	err = db.Setup()
	if err != nil {
		logrus.Fatal("Failed to set up database: ", err)
	}
	defer db.Close()

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

	// Setup logrus for fiber
	app.Use(fiberlog.New())

	logrus.Debug("adding favicon middleware")
	app.Use(favicon.New(favicon.Config{
		File: "./html/assets/favicon.ico",
	}))

	// Setup session store
	logrus.Debug("initializing session store")
	store := session.New(session.Config{
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:" + auth.SessionCookieName,
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

	// API routes
	apiGroup := app.Group("/api")
	v1 := apiGroup.Group("/v1")
	v1.Post("/users/add", api.AddUser())
	v1.Get("/users/get", api.GetUsers())
	v1.Delete("/users/delete", api.DeleteUser())
	v1.Get("/authenticated", api.Authenticated(store))

	logrus.Debug("registering static file server for assets")
	app.Static("/assets", "./html/assets")
	app.Static("/assets", "./admin-ui/assets")
	app.Static("/admin", "./admin-ui")

	// Start server
	logrus.Debug("starting web server")
	err = app.Listen(":80")
	if err != nil {
		logrus.Fatal("Failed to start web server: ", err)
	}
}
