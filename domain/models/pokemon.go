package models

import (
	"database/sql"
	"time"
)

type Pokemon struct {
	id        int64
	pokemonID int64
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func NewPokemon() *Pokemon {
	return &Pokemon{}
}

func (model *Pokemon) SetID(id int64) *Pokemon {
	model.id = id
	return model
}

func (model *Pokemon) SetPokemonID(pokemonID int64) *Pokemon {
	model.pokemonID = pokemonID
	return model
}

func (model *Pokemon) SetName(name string) *Pokemon {
	model.name = name
	return model
}

func (model *Pokemon) SetCreatedAt(createdAt time.Time) {
	model.createdAt = createdAt
}

func (model *Pokemon) SetUpdatedAt(updatedAt time.Time) {
	model.updatedAt = updatedAt
}

func (model *Pokemon) SetDeletedAt(deletedAt sql.NullTime) {
	model.deletedAt = deletedAt
}

func (model *Pokemon) ID() int64 {
	return model.id
}

func (model *Pokemon) PokemonID() int64 {
	return model.pokemonID
}

func (model *Pokemon) Name() string {
	return model.name
}

func (model *Pokemon) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Pokemon) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Pokemon) DeletedAt() sql.NullTime {
	return model.deletedAt
}

// Declare constant
const (
	pokemonSelectStatement    = `SELECT P.id, P.name, P.pokemon_id, P.created_at, P.updated_at, P.deleted_at FROM pokemon P`
	pokemonWhereStatement     = `WHERE LOWER(P.name) LIKE ? AND P.deleted_at IS NULL`
	pokemonDeletedAtStatement = `P.deleted_at IS NULL`
)

func (model Pokemon) GetSelect() string {
	return pokemonSelectStatement
}

func (model Pokemon) GetWhere() string {
	return pokemonWhereStatement
}

func (model Pokemon) GetDelete() string {
	return pokemonDeletedAtStatement
}

func (model *Pokemon) ScanRows(rows *sql.Rows) error {
	err := rows.Scan(&model.id, &model.name, &model.pokemonID, &model.createdAt, &model.updatedAt, &model.deletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (model *Pokemon) ScanRow(row *sql.Row) error {
	err := row.Scan(&model.id, &model.name, &model.pokemonID, &model.createdAt, &model.updatedAt, &model.deletedAt)
	if err != nil {
		return err
	}

	return nil
}
