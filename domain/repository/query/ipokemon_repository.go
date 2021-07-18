package query

import "github.com/faizalnurrozi/phincon-catch/domain/models"

type IPokemonRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []*models.Pokemon, err error)

	Count(search string) (res int, err error)

	CountBy(value map[string]interface{}) (res int, err error)
}

