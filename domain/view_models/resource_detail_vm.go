package view_models

type ResourceDetail struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
	Moves   []string `json:"moves"`
	Types   []string `json:"types"`
}
