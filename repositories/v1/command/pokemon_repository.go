package command

import (
	"database/sql"
	"github.com/faizalnurrozi/phincon-catch/domain/models"
	"github.com/faizalnurrozi/phincon-catch/domain/repository/command"
)

type PokemonRepository struct {
	DB *sql.DB
}

func NewPokemonRepository(DB *sql.DB) command.IPokemonRepository {
	return &PokemonRepository{DB: DB}
}

// Add data by request
func (repository PokemonRepository) Add(model *models.Pokemon, tx *sql.Tx) (res int64, err error) {
	statement := `INSERT INTO pokemon (pokemon_id, name, created_at, updated_at) VALUES (?,?,?,?)`
	result, err := tx.Exec(statement, model.PokemonID(), model.Name(), model.CreatedAt(), model.UpdatedAt())
	if err != nil {
		return res, err
	}

	res, _ = result.LastInsertId()

	return res, nil
}

// Edit data by request
func (repository PokemonRepository) Edit(model *models.Pokemon, tx *sql.Tx) (res int64, err error) {
	statement := `UPDATE pokemon SET name = ?, updated_at = ? WHERE id = ?`
	result, err := tx.Exec(statement, model.Name(), model.UpdatedAt(), model.ID())
	if err != nil {
		return res, err
	}

	res, _ = result.LastInsertId()

	return res, nil
}

// DeleteBy data by column
func (repository PokemonRepository) DeleteBy(column, value, operator string, model *models.Pokemon, tx *sql.Tx) (err error) {
	statement := `UPDATE pokemon SET updated_at = ?, deleted_at = ? WHERE ` + column + ` ` + operator + ` ?`
	_, err = tx.Exec(statement, model.UpdatedAt(), model.DeletedAt().Time, value)
	if err != nil {
		return err
	}

	return nil
}
