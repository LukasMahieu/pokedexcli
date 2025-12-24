package main

import (
	"errors"
	"fmt"
	"github.com/LukasMahieu/pokedexcli/internal/pokeapi"
	"github.com/LukasMahieu/pokedexcli/internal/pokecache"
	"os"
)

type config struct {
	Previous string
	Next     string
	Cache    pokecache.Cache
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
