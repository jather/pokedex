package main

import (
	"time"

	"github.com/jather/pokedexcli/internal/pokeapi"
	"github.com/jather/pokedexcli/internal/pokecache"
)

const prompt = "Pokedex > "

func main() {
	pokeapiclient := pokeapi.NewHttpClient(time.Second * 5)
	cfg := &config{httpClient: pokeapiclient}
	cache := pokecache.NewCache(10 * time.Second)
	cfg.cache = &cache
	startRepl(cfg)
}
