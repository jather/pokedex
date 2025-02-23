package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, arg string) error {
	if arg == "" {
		return errors.New("must specify area")
	}
	client := cfg.httpClient
	response, err := client.GetAreaPokemons(arg)
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
