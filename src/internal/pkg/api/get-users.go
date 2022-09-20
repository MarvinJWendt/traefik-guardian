package api

import (
	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func GetUsers() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		users, err := db.GetUsers()
		if err != nil {
			return err
		}

		return ctx.JSON(users)
	}
}
