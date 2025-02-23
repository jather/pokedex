package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arg string) error {
	if arg == "" {
		return errors.New("must specify pokemon")
	}
	client := cfg.httpClient
	response, err := client.GetPokemon(arg)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	catchChance := (0.5 - 0.0015*float64(response.BaseExperience)) + 0.3
	if rand.Float64() > catchChance {
		fmt.Printf("%s got away!\n", arg)
		return nil
	}
	fmt.Printf("%s was caught!\n", arg)
	cfg.pokedex[arg] = response
	return nil
}
