package main

import "fmt"

func commandPokedex(cfg *config, _ string) error {
	fmt.Println("Your pokemon:")
	for pokemon := range cfg.pokedex {
		fmt.Println(" -" + pokemon)
	}
	return nil
}
