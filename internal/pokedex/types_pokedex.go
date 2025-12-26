package pokedex

import (
	"github.com/LukasMahieu/pokedexcli/internal/pokeapi"
)

type Pokedex struct {
	pokemons map[string]pokeapi.Pokemon
}
