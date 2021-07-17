package v1

import (
	"database/sql"
	"github.com/faizalnurrozi/phincon-catch/domain/models"
	"github.com/faizalnurrozi/phincon-catch/domain/request"
	"github.com/faizalnurrozi/phincon-catch/domain/usecases"
	"github.com/faizalnurrozi/phincon-catch/packages/functioncaller"
	"github.com/faizalnurrozi/phincon-catch/packages/logruslogger"
	"github.com/faizalnurrozi/phincon-catch/repositories/v1/command"
	"github.com/faizalnurrozi/phincon-catch/repositories/v1/query"
	"github.com/faizalnurrozi/phincon-catch/usecase"
	"time"
)

type PokemonUseCase struct {
	*usecase.Contract
}

func NewPokemonUseCase(ucContract *usecase.Contract) usecases.IPokemonUseCase {
	return &PokemonUseCase{Contract: ucContract}
}

// Add data by request
func (uc PokemonUseCase) Add(req *request.PokemonRequest) (res int64, err error) {

	// Check if name is already exists
	payloads := map[string]interface{}{"name": req.Name}
	_, err = uc.countBy(payloads)

	// Handle response and error if name is already exists
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-pokemon-countBy")
		return res, err
	}

	// Initiate repository
	repository := command.NewPokemonRepository(uc.DB)
	now := time.Now().UTC()

	// Initiate model with request data
	model := models.NewPokemon()
	model.SetName(req.Name)
	model.SetPokemonID(req.PokemonID)
	model.SetCreatedAt(now)
	model.SetUpdatedAt(now)

	// process model to database
	res, err = repository.Add(model, uc.TX)

	// Handle response after query from database
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-pokemon-add")
		return res, err
	}

	return res, nil
}

// Edit data by request
func (uc PokemonUseCase) Edit(req *request.PokemonRequest, ID int64) (res int64, err error) {

	// Check if name is already exists
	payloads := map[string]interface{}{"name": req.Name, "id": ID}
	_, err = uc.countBy(payloads)

	// Handle response and error if name is already exists
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-pokemon-countBy")
		return res, err
	}

	// Initiate repository
	repository := command.NewPokemonRepository(uc.DB)
	now := time.Now().UTC()

	// Initiate model with request data
	model := models.NewPokemon()
	model.SetID(ID)
	model.SetPokemonID(req.PokemonID)
	model.SetName(req.Name)
	model.SetUpdatedAt(now)

	// Process model to database
	res, err = repository.Edit(model, uc.TX)

	// Handle response after query from database
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-pokemon-edit")
		return res, err
	}

	return res, nil
}

// Delete data by column
func (uc PokemonUseCase) DeleteBy(column, value, operator string) (err error) {

	// Initiate repository
	repository := command.NewPokemonRepository(uc.DB)
	now := time.Now().UTC()

	// Process model to database
	model := models.NewPokemon()
	model.SetUpdatedAt(now)
	model.SetDeletedAt(sql.NullTime{Time: now, Valid: true})
	err = repository.DeleteBy(column, value, operator, model, uc.TX)

	// Handle response after query from database
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-pokemon-delete")
		return err
	}

	return nil
}

// Function for get count of data
func (uc PokemonUseCase) Count(search string) (res int, err error) {

	// Initiate repository
	repository := query.NewPokemonRepository(uc.DB)

	// Query count data from database
	res, err = repository.Count(search)

	// Handle response after query from database
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-pokemon-count")
		return res, err
	}

	return res, nil
}

func (uc PokemonUseCase) countBy(payloads map[string]interface{}) (res int, err error) {

	// Initiate repository
	repository := query.NewPokemonRepository(uc.DB)

	// Query count data
	res, err = repository.CountBy(payloads)

	// Response from repository
	if err != nil {
		return res, err
	}

	return res, err
}
