package handlers

import (
	"database/sql"
	"github.com/faizalnurrozi/phincon-catch/usecase"
	"net/http"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/faizalnurrozi/phincon-catch/packages/str"
	"github.com/faizalnurrozi/phincon-catch/domain/view_models"
)

type Handler struct {
	App        *fiber.App
	UcContract *usecase.Contract
	DB         *sql.DB
	Validate   *validator.Validate
	Translator ut.Translator
}

const (
	ResponseWithMeta              = `responseWithMeta`
	ResponseWithOutMeta           = `responseWithOutMeta`
	ResponseErrorValidationStruct = `responseErrorValidationStruct`
)

//send response
func (handler Handler) SendResponse(ctx *fiber.Ctx, responseType string, data interface{}, meta interface{}, err interface{}, statusCode int) error {

	if err != nil {
		if statusCode == 400 || responseType == ResponseErrorValidationStruct {
			errorMessage := err.(validator.ValidationErrors)
			return handler.ResponseValidationError(ctx, errorMessage)
		} else {
			errorMessage := err.(error)
			return handler.SendErrorResponse(ctx, errorMessage.Error(), statusCode)
		}
	} else {
		if responseType == ResponseWithMeta {
			return handler.SendSuccessResponseWithMeta(ctx, data, meta)
		} else {
			return handler.SendSuccessResponseWithOutMeta(ctx, data)
		}
	}
}

//success response with out meta
func (handler Handler) SendSuccessResponseWithOutMeta(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(view_models.SuccessResponseWithOutMetaVm{Data: data})
}

//success response with meta
func (handler Handler) SendSuccessResponseWithMeta(ctx *fiber.Ctx, data interface{}, meta interface{}) error {
	return ctx.Status(http.StatusOK).JSON(view_models.SuccessResponseWithMeta{Data: data, Meta: meta})
}

//error response
func (handler Handler) SendErrorResponse(ctx *fiber.Ctx, err interface{}, statusCode int) error {
	return ctx.Status(statusCode).JSON(view_models.ErrorResponseVm{Message: err})
}

//validation error
func (handler Handler) ResponseValidationError(ctx *fiber.Ctx, error validator.ValidationErrors) error {
	errorMessage := handler.ExtractErrorValidationMessages(error)

	return handler.SendErrorResponse(ctx, errorMessage, http.StatusBadRequest)
}

func (handler Handler) ExtractErrorValidationMessages(error validator.ValidationErrors) map[string][]string {
	errorMessage := map[string][]string{}
	errorTranslation := error.Translate(handler.Translator)

	for _, err := range error {
		errKey := str.Underscore(err.StructField())
		errorMessage[errKey] = append(
			errorMessage[errKey],
			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), err.StructField(), -1),
		)
	}

	return errorMessage
}

func (handler Handler) StringToLowerAndRemoveWhiteSpaces(text string) string {
	return strings.TrimSpace(strings.ToLower(text))
}

func (handler Handler) ArrayStringToLowerAndRemoveWhiteSpaces(arrayText []string) []string {
	for i, val := range arrayText {
		arrayText[i] = handler.StringToLowerAndRemoveWhiteSpaces(val)
	}

	return arrayText
}

func (handler Handler) ArrayStringToLowerAndRemoveWhiteSpacesAndJoinBack(arrayText []string) string {
	arrayText = handler.ArrayStringToLowerAndRemoveWhiteSpaces(arrayText)

	return strings.Join(arrayText, ",")
}
