package v1

import (
	"database/sql"
	"github.com/faizalnurrozi/phincon-catch/domain/models"
	"github.com/faizalnurrozi/phincon-catch/domain/request"
	"github.com/faizalnurrozi/phincon-catch/domain/usecases"
	"github.com/faizalnurrozi/phincon-catch/domain/view_models"
	"github.com/faizalnurrozi/phincon-catch/packages/functioncaller"
	"github.com/faizalnurrozi/phincon-catch/packages/logruslogger"
	"github.com/faizalnurrozi/phincon-catch/repositories/v1/command"
	"github.com/faizalnurrozi/phincon-catch/repositories/v1/query"
	"github.com/faizalnurrozi/phincon-catch/usecase"
	"math/rand"
	"time"
)

type PokemonUseCase struct {
	*usecase.Contract
}

func NewPokemonUseCase(ucContract *usecase.Contract) usecases.IPokemonUseCase {
	return &PokemonUseCase{Contract: ucContract}
}

func (uc PokemonUseCase) randomProbability() bool {
	return rand.Float32() < 0.5
}

// Browse all data by ordering and sorting
func (uc PokemonUseCase) Browse(search, orderBy, sort string, page, limit int) (res []view_models.PokemonVM, pagination view_models.PaginationVm, err error) {

	// Initiate repository and view model
	repository := query.NewPokemonRepository(uc.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)
	vm := view_models.PokemonVM{}

	// Query data from database
	pokemons, err := repository.Browse(search, orderBy, sort, limit, offset)

	// Handle response after query from database
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-pokemon-browse")
		return res, pagination, err
	}

	// Append data to view model
	for _, pokemon := range pokemons {
		res = append(res, vm.Build(pokemon))
	}

	//set pagination
	totalCount, err := uc.Count(search)

	// Handle response after query from database
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-pokemon-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

// Catch data by 0.5 probability
func (uc PokemonUseCase) Catch(req *request.PokemonCatchRequest) (res view_models.PokemonCatch, err error) {
	res = view_models.PokemonCatch{
		PokemonID: req.PokemonID,
		Status:    uc.randomProbability(),
	}

	return res, err
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
