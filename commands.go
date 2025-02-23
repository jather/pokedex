package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jather/pokedexcli/internal/pokeapi"
	"github.com/jather/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}
type config struct {
	httpClient pokeapi.Client
	previous   *string
	next       *string
	cache      *pokecache.Cache
}

func getCommands() map[string]cliCommand {
	commandList := map[string]cliCommand{
		"exit": {
			name:        "Exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "Help",
			description: "See available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "Map",
			description: "See a list of 20 locations. Use the command again to see the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "Mapb",
			description: "See the previous page from the list of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "Explore",
			description: "Explore an area and see the list of pokemon that can be found. Syntax: explore <area>",
			callback:    commandExplore,
		},
	}
	return commandList
}

func commandExit(cfg *config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("did not exit properly")
}
func commandHelp(cfg *config, _ string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Println(command.name, ": ", command.description)
	}
	return nil
}
func commandMap(cfg *config, _ string) error {
	client := cfg.httpClient
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		url = *cfg.next
	}

	response, err := client.GetAreas(url, cfg.cache)
	if err != nil {
		return err
	}
	cfg.next = response.Next
	cfg.previous = response.Previous
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, _ string) error {
	client := cfg.httpClient
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}
	response, err := client.GetAreas(*cfg.previous, cfg.cache)
	if err != nil {
		return err
	}
	cfg.next = response.Next
	cfg.previous = response.Previous
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	return nil
}
func commandExplore(cfg *config, arg string) error {
	if arg == "" {
		return errors.New("must specify area")
	}
	client := cfg.httpClient
	response, err := client.GetAreaPokemons(arg, cfg.cache)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", arg)
	fmt.Println("Found Pokemon:")
	for _, encounter := range response.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
