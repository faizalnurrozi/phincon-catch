package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/faizalnurrozi/phincon-catch/packages/messages"
	"github.com/faizalnurrozi/phincon-catch/server/http/handlers"
	"github.com/faizalnurrozi/phincon-catch/usecase"
	"net/http"
	"os"
)

type BasicAuth struct {
	*usecase.Contract
}

func (basicAuth BasicAuth) BasicAuthNew() fiber.Handler {
	handler := handlers.Handler{UcContract: basicAuth.Contract}
	configAuth := basicauth.Config{
		Users: map[string]string{
			os.Getenv("BASIC_AUTH_USERNAME_CLIENT"): os.Getenv("BASIC_AUTH_PASSWORD_CLIENT"),
		},
		Realm: "Forbidden",
		Unauthorized: func(ctx *fiber.Ctx) error {
			return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}
	return basicauth.New(configAuth)
}
