package api

import (
	"strconv"

	"github.com/MarvinJWendt/traefik-auth-provider/src/internal/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func DeleteUser() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		username := ctx.Query("username")
		id := ctx.Query("id")
		var idInt int

		if username == "" && id == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("username and id are empty - one of them is needed")
		}

		// validate that ID is a number if it's set
		if id != "" {
			i, err := strconv.Atoi(id)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString("id is not a number")
			}
			idInt = i
		}

		var user db.User

		if idInt != 0 {
			u, err := db.GetUserByID(idInt)
			if err != nil {
				return err
			}
			user = u
		} else {
			u, err := db.GetUserByUsername(username)
			if err != nil {
				return err
			}
			user = u
		}

		err := db.DeleteUser(user.ID)
		if err != nil {
			return err
		}

		return nil
	}
}
