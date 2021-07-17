package query

import (
	"database/sql"
	"errors"
	"github.com/faizalnurrozi/phincon-catch/domain/models"
	"github.com/faizalnurrozi/phincon-catch/domain/repository/query"
	"github.com/faizalnurrozi/phincon-catch/packages/messages"
	"strings"
)

type PokemonRepository struct {
	DB *sql.DB
}

func NewPokemonRepository(DB *sql.DB) query.IPokemonRepository {
	return &PokemonRepository{DB: DB}
}

// Count Get count data from table
func (repository PokemonRepository) Count(search string) (res int, err error) {
	model := models.NewPokemon()
	statement := `SELECT COUNT(P.id) FROM pokemon P ` + model.GetWhere()
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository PokemonRepository) CountBy(payloads map[string]interface{}) (res int, err error) {
	var whereStatement string
	whereStatement += "WHERE P.deleted_at IS NULL AND LOWER(P.name) = ?"
	whereParams := []interface{}{strings.ToLower(payloads["name"].(string))}

	if val, ok := payloads["id"]; ok {
		whereStatement += " AND P.id <> ?"
		whereParams = append(whereParams, val.(int64))
	}

	statement := `SELECT COUNT(P.id) FROM pokemon P ` + whereStatement
	err = repository.DB.QueryRow(statement, whereParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	if res > 0 {
		return res, errors.New(messages.DataAlreadyExist)
	}

	return res, err
}
