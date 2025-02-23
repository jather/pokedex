package pokeapi

import (
	"net/http"
	"time"

	"github.com/jather/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewHttpClient(timeout time.Duration, cache *pokecache.Cache) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      cache,
	}
}
