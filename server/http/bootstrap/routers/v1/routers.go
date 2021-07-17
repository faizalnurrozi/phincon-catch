package v1

import (
	"github.com/faizalnurrozi/phincon-catch/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type Routers struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (routers Routers) RegisterRoute() {
	apiV1 := routers.RouteGroup.Group("/v1")

	pokemonRoutes := PokemonRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	pokemonRoutes.RegisterRoute()
}
