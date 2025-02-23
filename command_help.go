package main

import "fmt"

func commandHelp(cfg *config, _ string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Println(command.name, ": ", command.description)
	}
	return nil
}
