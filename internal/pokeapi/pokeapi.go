package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/jather/pokedexcli/internal/pokecache"
)

type areaLocationResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetAreas(url string, cache *pokecache.Cache) (areaLocationResponse, error) {
	// sends http request to get area, returns, arealocationresponse struct

	// check if url is in cache
	val, ok := cache.Get(url)
	if ok {
		areas := areaLocationResponse{}
		if err := json.Unmarshal(val, &areas); err != nil {
			return areaLocationResponse{}, errors.New("error while decoding to json data to struct")
		}
		fmt.Println("from cached:")
		return areas, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return areaLocationResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return areaLocationResponse{}, nil
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return areaLocationResponse{}, err
	}
	cache.Add(url, data)
	areas := areaLocationResponse{}
	if err := json.Unmarshal(data, &areas); err != nil {
		return areaLocationResponse{}, errors.New("error while decoding to json data to struct")
	}
	return areas, nil

}
