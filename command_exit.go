package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("did not exit properly")
}
