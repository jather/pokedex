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
type areaLocationDetailedResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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
func (c *Client) GetAreaPokemons(area string, cache *pokecache.Cache) (areaLocationDetailedResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + area
	val, ok := cache.Get(url)
	if ok {
		areaData := areaLocationDetailedResponse{}
		if err := json.Unmarshal(val, &areaData); err != nil {
			return areaLocationDetailedResponse{}, err
		}
		fmt.Println("from cached:")
		return areaData, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return areaLocationDetailedResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return areaLocationDetailedResponse{}, nil
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return areaLocationDetailedResponse{}, nil
	}
	if res.StatusCode > 299 {
		return areaLocationDetailedResponse{}, errors.New("unsucessful request. check if the area you specified is correct")
	}
	areaData := areaLocationDetailedResponse{}
	err = json.Unmarshal(data, &areaData)
	if err != nil {
		return areaLocationDetailedResponse{}, err
	}
	return areaData, nil

}
