package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jather/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
type config struct {
	httpClient pokeapi.Client
	previous   *string
	next       *string
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
	}
	return commandList
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("did not exit properly")
}
func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Println(command.name, ": ", command.description)
	}
	return nil
}
func commandMap(cfg *config) error {
	client := cfg.httpClient
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		url = *cfg.next
	}
	response, err := client.GetAreas(url)
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

func commandMapb(cfg *config) error {
	client := cfg.httpClient
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}
	response, err := client.GetAreas(*cfg.previous)
	if err != nil {
		return err
	}
	fmt.Println(response)
	cfg.next = response.Next
	cfg.previous = response.Previous
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	return nil
}
