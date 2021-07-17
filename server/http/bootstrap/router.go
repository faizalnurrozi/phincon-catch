package bootstrap

import (
	"github.com/faizalnurrozi/phincon-catch/server/http/bootstrap/routers/v1"
	"github.com/faizalnurrozi/phincon-catch/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func (boot Bootstrap) RegisterRoute() {
	handlerType := handlers.Handler{
		App:        boot.App,
		UcContract: &boot.UcContract,
		DB:         boot.Db,
		Validate:   boot.Validator,
		Translator: boot.Translator,
	}

	// Route for check health
	rootParentGroup := boot.App.Group("/phincon-browse")
	rootParentGroup.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("Work")
	})

	// Grouping v1 api
	v1Routers := v1.Routers{
		RouteGroup: rootParentGroup,
		Handler:    handlerType,
	}
	v1Routers.RegisterRoute()
}
