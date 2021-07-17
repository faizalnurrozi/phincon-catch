package v1

import (
	"github.com/faizalnurrozi/phincon-catch/server/http/handlers"
	v1 "github.com/faizalnurrozi/phincon-catch/server/http/handlers/v1"
	"github.com/faizalnurrozi/phincon-catch/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type PokemonRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route PokemonRoute) RegisterRoute() {

	// Initiate handler
	handler := v1.PokemonHandler{Handler: route.Handler}
	basicAuthMiddleware := middlewares.BasicAuth{Contract: route.Handler.UcContract}

	// List of route
	pokemonGroupRouters := route.RouteGroup.Group("/pokemon")
	pokemonGroupRouters.Use(basicAuthMiddleware.BasicAuthNew())
	pokemonGroupRouters.Post("/catch", handler.Catch)
	pokemonGroupRouters.Post("", handler.Add)
	pokemonGroupRouters.Put("/:id", handler.Edit)
	pokemonGroupRouters.Delete("/:id", handler.DeleteBy)
}
