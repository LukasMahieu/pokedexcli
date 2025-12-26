package pokedex

import (
	"fmt"
	"github.com/LukasMahieu/pokedexcli/internal/pokeapi"
)

func NewPokedex() *Pokedex {
	return &Pokedex{
		pokemons: make(map[string]pokeapi.Pokemon),
	}
}

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) {
	p.pokemons[pokemon.Name] = pokemon
}

func (p *Pokedex) ListPokemon() {
	fmt.Println("Your Pokedex:")
	for name := range p.pokemons {
		fmt.Printf("  - %s\n", name)
	}
}

func (p *Pokedex) Inspect(pokemonName string) {
	pokemon, ok := p.pokemons[pokemonName]
	if !ok {
		fmt.Println("this pokemon has not been caught yet")
		return
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, stat := range pokemon.Types {
		fmt.Printf("  -%s\n", stat.Type.Name)
	}
}
