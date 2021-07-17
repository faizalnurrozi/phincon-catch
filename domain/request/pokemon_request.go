package request

type PokemonRequest struct {
	PokemonID int64  `json:"pokemon_id"`
	Name      string `json:"name"`
}
