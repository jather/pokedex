package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewHttpClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
	}
}
