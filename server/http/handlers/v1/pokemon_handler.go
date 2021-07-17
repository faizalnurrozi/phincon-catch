package v1

import (
	handlers2 "github.com/faizalnurrozi/phincon-catch/domain/handlers"
	"github.com/faizalnurrozi/phincon-catch/domain/request"
	"github.com/faizalnurrozi/phincon-catch/server/http/handlers"
	v1 "github.com/faizalnurrozi/phincon-catch/usecase/v1"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type PokemonHandler struct {
	handlers.Handler
}

func NewPokemonHandler(handler handlers.Handler) handlers2.IPokemonHandler {
	return &PokemonHandler{Handler: handler}
}

// function handler for add data by payload
func (handler PokemonHandler) Catch(ctx *fiber.Ctx) (err error) {

	// Parse & Checking input
	input := new(request.PokemonCatchRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}

	// Database processing
	handler.UcContract.TX, err = handler.DB.Begin()
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	uc := v1.NewPokemonUseCase(handler.UcContract)
	res, err := uc.Catch(input)
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

// function handler for add data by payload
func (handler PokemonHandler) Add(ctx *fiber.Ctx) (err error) {

	// Parse & Checking input
	input := new(request.PokemonRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}

	// Database processing
	handler.UcContract.TX, err = handler.DB.Begin()
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	uc := v1.NewPokemonUseCase(handler.UcContract)
	res, err := uc.Add(input)
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

// function handler for edit data by payload
func (handler PokemonHandler) Edit(ctx *fiber.Ctx) (err error) {
	// Get Param
	ID := ctx.Params("id")

	// Parse & Checking input
	input := new(request.PokemonRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}

	// Database processing
	handler.UcContract.TX, err = handler.DB.Begin()
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	uc := v1.NewPokemonUseCase(handler.UcContract)
	idConverted, _ := strconv.Atoi(ID)
	res, err := uc.Edit(input, int64(idConverted))
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

// function handler for delete data by ID
func (handler PokemonHandler) DeleteBy(ctx *fiber.Ctx) (err error) {
	// Get Param
	ID := ctx.Params("id")

	// Database processing
	handler.UcContract.TX, err = handler.DB.Begin()
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}
	uc := v1.NewPokemonUseCase(handler.UcContract)
	err = uc.DeleteBy("id", ID, "=")
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
}
