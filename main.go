package main

import (
	"time"

	"github.com/jather/pokedexcli/internal/pokeapi"
	"github.com/jather/pokedexcli/internal/pokecache"
)

const prompt = "Pokedex > "

func main() {
	cache := pokecache.NewCache(10 * time.Second)
	pokeapiclient := pokeapi.NewHttpClient(time.Second*5, &cache)
	cfg := &config{httpClient: pokeapiclient, pokedex: map[string]pokeapi.PokemonResponse{}}
	startRepl(cfg)
}
