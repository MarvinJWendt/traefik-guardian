package api

import (
	"errors"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/db"
	"github.com/asdine/storm/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func AddUser() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		username := ctx.Query("username")
		password := ctx.Query("password")

		if username == "" || password == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("username or password is empty")
		}

		err := db.CreateUser(db.User{
			Username: username,
			Password: password,
		})

		if err != nil {
			if errors.Is(err, storm.ErrAlreadyExists) {
				return ctx.Status(fiber.StatusConflict).SendString("user already exists")
			}

			logrus.Error(err)

			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return nil
	}
}
