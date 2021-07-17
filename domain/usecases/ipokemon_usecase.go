package usecases

import (
	"github.com/faizalnurrozi/phincon-catch/domain/request"
)

type IPokemonUseCase interface {
	Add(req *request.PokemonRequest) (res int64, err error)

	Edit(req *request.PokemonRequest, ID int64) (res int64, err error)

	DeleteBy(column, value, operator string) (err error)

	Count(search string) (res int, err error)
}
