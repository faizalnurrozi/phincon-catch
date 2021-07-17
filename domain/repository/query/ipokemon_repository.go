package query

type IPokemonRepository interface {
	Count(search string) (res int, err error)

	CountBy(value map[string]interface{}) (res int, err error)
}

