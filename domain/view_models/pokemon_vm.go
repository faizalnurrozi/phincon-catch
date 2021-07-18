package view_models

import (
	"github.com/faizalnurrozi/phincon-catch/domain/models"
	"github.com/faizalnurrozi/phincon-catch/packages/pokeapi-go"
	"strconv"
)

type PokemonVM struct {
	ID        int64  `json:"id"`
	PokemonID int64  `json:"pokemon_id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
}

// Build model for this struct
func (vm PokemonVM) Build(model *models.Pokemon) PokemonVM {

	pokemonPict, _ := pokeapi.Pokemon(strconv.Itoa(int(model.PokemonID())))

	return PokemonVM{
		ID:        model.ID(),
		PokemonID: model.PokemonID(),
		Name:      model.Name(),
		Picture:   pokemonPict.Sprites.FrontDefault,
	}
}
