package main

import (
	"time"

	"github.com/jather/pokedexcli/internal/pokeapi"
)

const prompt = "Pokedex > "

func main() {
	pokeapiclient := pokeapi.NewHttpClient(time.Second * 5)
	cfg := &config{httpClient: pokeapiclient}
	startRepl(cfg)
}
