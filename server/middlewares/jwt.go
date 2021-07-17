package middlewares

import (
	"errors"
	"fmt"
	"github.com/faizalnurrozi/phincon-catch/server/http/handlers"
	"github.com/faizalnurrozi/phincon-catch/usecase"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/faizalnurrozi/phincon-catch/packages/functioncaller"
	jwtPkg "github.com/faizalnurrozi/phincon-catch/packages/jwt"
	"github.com/faizalnurrozi/phincon-catch/packages/logruslogger"
	"github.com/faizalnurrozi/phincon-catch/packages/messages"
)

type JwtMiddleware struct {
	*usecase.Contract
}

func (jwtMiddleware JwtMiddleware) New(ctx *fiber.Ctx) (err error) {
	claims := &jwtPkg.CustomClaims{}
	handler := handlers.Handler{UcContract: jwtMiddleware.Contract}

	//check header is present or not
	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "middleware-jwt-checkHeader")
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//check claims and signing method
	token := strings.Replace(header, "Bearer ", "", -1)
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			logruslogger.Log(logruslogger.WarnLevel, messages.UnexpectedSigningMethod, functioncaller.PrintFuncName(), "middleware-jwt-checkSigningMethod")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := []byte(jwtMiddleware.JwtCredential.TokenSecret)
		return secret, nil
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "middleware-jwt-checkClaims")
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//check token live time
	if claims.ExpiresAt < time.Now().Unix() {
		logruslogger.Log(logruslogger.WarnLevel, messages.ExpiredToken, functioncaller.PrintFuncName(), "middleware-jwt-checkTokenLiveTime")
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//jwe roll back encrypted id
	jweRes, err := jwtMiddleware.JweCredential.Rollback(claims.Payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwe-rollback")
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}
	if jweRes == nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "pkg-jwe-resultNil")
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//set id to uce case contract
	claims.Id = fmt.Sprintf("%v", jweRes["id"])
	jwtMiddleware.Contract.UserID = claims.Id
	jwtMiddleware.Contract.RoleID = fmt.Sprintf("%v", jweRes["roleID"])

	return ctx.Next()
}
