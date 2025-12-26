package main

import (
	"errors"
	"fmt"
	"github.com/LukasMahieu/pokedexcli/internal/pokeapi"
	"github.com/LukasMahieu/pokedexcli/internal/pokecache"
	"github.com/LukasMahieu/pokedexcli/internal/pokedex"
	"math/rand"
	"os"
	"time"
)

type config struct {
	Previous string
	Next     string
	Cache    pokecache.Cache
	Pokedex  *pokedex.Pokedex
}

func catchChance(baseExperience, k float64) float64 {
	normalized := baseExperience / (baseExperience + k)
	return 1 - normalized
}

func tryCatch(baseChance float64) bool {
	return rand.Float64() < baseChance
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Failed to exit")
}

func commandHelp(cfg *config, args []string) error {
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config, args []string) error {
	url := cfg.Next
	locations, err := pokeapi.FetchLocations(url, cfg.Cache)
	if err != nil {
		return err
	}
	cfg.Previous = locations.Previous
	cfg.Next = locations.Next
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapBack(cfg *config, args []string) error {
	url := cfg.Previous

	locations, err := pokeapi.FetchLocations(url, cfg.Cache)
	if err != nil {
		return err
	}
	cfg.Previous = locations.Previous
	cfg.Next = locations.Next
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a location name")
	}
	location := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + location
	result, err := pokeapi.FetchSpecificLocation(url, cfg.Cache)
	if err != nil {
		return err
	}
	for _, encounter := range result.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a pokemon's name to catch")
	}
	if args[0] == " " {
		return errors.New("please provide a pokemon's name to catch")
	}
	pokemon := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	result, err := pokeapi.FetchPokemonInfo(url, cfg.Cache)
	if err != nil {
		return err
	}

	// higher base xp --> more difficult to catch pokemon
	baseExperience := float64(result.BaseExperience) // value between [0,+inf[
	catchChance := catchChance(baseExperience, 100)
	caught := tryCatch(catchChance)

	// interface
	fmt.Printf("Throwing a Pokeball at %s...\n", result.Name)
	time.Sleep(1 * time.Second)
	if caught {
		fmt.Printf("%s was caught!\n", result.Name)
		cfg.Pokedex.Add(result)
	} else {
		fmt.Printf("%s escaped!\n", result.Name)
	}
	return nil
}

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a pokemon's name to catch")
	}
	if args[0] == " " {
		return errors.New("please provide a pokemon's name to catch")
	}
	pokemon := args[0]
	cfg.Pokedex.Inspect(pokemon)
	return nil
}

func commandPokedex(cfg *config, args []string) error {
	cfg.Pokedex.ListPokemon()
	return nil
}
