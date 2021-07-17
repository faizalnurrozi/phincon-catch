package command

import (
	"database/sql"
	"github.com/faizalnurrozi/phincon-catch/domain/models"
)

type IPokemonRepository interface {
	Add(model *models.Pokemon, tx *sql.Tx) (res int64, err error)

	Edit(model *models.Pokemon, tx *sql.Tx) (res int64, err error)

	DeleteBy(column, value, operator string, model *models.Pokemon, tx *sql.Tx) (err error)
}

