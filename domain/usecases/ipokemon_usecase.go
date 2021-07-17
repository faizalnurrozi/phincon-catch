package usecases

import (
	"github.com/faizalnurrozi/phincon-catch/domain/request"
	"github.com/faizalnurrozi/phincon-catch/domain/view_models"
)

type IPokemonUseCase interface {
	Catch(req *request.PokemonCatchRequest) (res view_models.PokemonCatch, err error)

	Add(req *request.PokemonRequest) (res int64, err error)

	Edit(req *request.PokemonRequest, ID int64) (res int64, err error)

	DeleteBy(column, value, operator string) (err error)

	Count(search string) (res int, err error)
}
