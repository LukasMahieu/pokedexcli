package main

import (
	"errors"
	"fmt"
	"github.com/LukasMahieu/pokedexcli/internal/pokeapi"
	"os"
)

type config struct {
	Previous string
	Next     string
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Failed to exit")
}

func commandHelp(cfg *config) error {
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	url := cfg.Next

	locations, err := pokeapi.FetchLocations(url)
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

func commandMapBack(cfg *config) error {
	url := cfg.Previous

	locations, err := pokeapi.FetchLocations(url)
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
