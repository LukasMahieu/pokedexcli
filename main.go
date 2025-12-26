package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/LukasMahieu/pokedexcli/internal/pokecache"
	"github.com/LukasMahieu/pokedexcli/internal/pokedex"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next set of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Displays the next set of locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Pokedex a pokemon",
			callback:    commandPokedex,
		},
	}
	return commands
}

func callCommand(cfg *config, command string, args []string) error {
	cmd, ok := getCommands()[command]
	if !ok {
		fmt.Println("Unknown command")
		return nil
	}
	return cmd.callback(cfg, args)

}

func main() {
	cache := pokecache.NewCache(time.Second * 5)

	cfg := &config{
		Previous: "",
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Cache:    cache,
		Pokedex:  pokedex.NewPokedex(),
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for i := 0; ; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		callCommand(cfg, cleanedInput[0], cleanedInput[1:])
	}
}
