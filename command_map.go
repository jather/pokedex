package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, _ string) error {
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

func commandMapb(cfg *config, _ string) error {
	client := cfg.httpClient
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}
	response, err := client.GetAreas(*cfg.previous)
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
