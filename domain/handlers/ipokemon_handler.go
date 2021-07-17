package handlers

import "github.com/gofiber/fiber/v2"

type IPokemonHandler interface {
	Catch(ctx *fiber.Ctx) (err error)

	Add(ctx *fiber.Ctx) (err error)

	Edit(ctx *fiber.Ctx) (err error)

	DeleteBy(ctx *fiber.Ctx) (err error)
}
