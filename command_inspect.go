package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokemon string) error {
	val, ok := cfg.pokedex[pokemon]
	if !ok {
		return errors.New("you don't have this pokemon")
	}
	fmt.Printf("Name: %s\nHeight: %v\nWeight: %v\n Stats:\n", val.Name, val.Height, val.Weight)
	for _, stat := range val.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range val.Types {
		fmt.Printf("  - %v\n", pokemonType.Type.Name)
	}
	return nil
}
