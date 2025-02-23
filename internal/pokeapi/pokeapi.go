package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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
type PokemonResponse struct {
	// Abilities []struct {
	// 	Ability struct {
	// 		Name string `json:"name"`
	// 		URL  string `json:"url"`
	// 	} `json:"ability"`
	// 	IsHidden bool `json:"is_hidden"`
	// 	Slot     int  `json:"slot"`
	// } `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	// Cries          struct {
	// 	Latest string `json:"latest"`
	// 	Legacy string `json:"legacy"`
	// } `json:"cries"`
	// Forms []struct {
	// 	Name string `json:"name"`
	// 	URL  string `json:"url"`
	// } `json:"forms"`
	// GameIndices []struct {
	// 	GameIndex int `json:"game_index"`
	// 	Version   struct {
	// 		Name string `json:"name"`
	// 		URL  string `json:"url"`
	// 	} `json:"version"`
	// } `json:"game_indices"`
	Height int `json:"height"`
	// HeldItems []struct {
	// 	Item struct {
	// 		Name string `json:"name"`
	// 		URL  string `json:"url"`
	// 	} `json:"item"`
	// 	VersionDetails []struct {
	// 		Rarity  int `json:"rarity"`
	// 		Version struct {
	// 			Name string `json:"name"`
	// 			URL  string `json:"url"`
	// 		} `json:"version"`
	// 	} `json:"version_details"`
	// } `json:"held_items"`
	ID int `json:"id"`
	// IsDefault              bool   `json:"is_default"`
	// LocationAreaEncounters string `json:"location_area_encounters"`
	// Moves                  []struct {
	// 	Move struct {
	// 		Name string `json:"name"`
	// 		URL  string `json:"url"`
	// 	} `json:"move"`
	// 	VersionGroupDetails []struct {
	// 		LevelLearnedAt  int `json:"level_learned_at"`
	// 		MoveLearnMethod struct {
	// 			Name string `json:"name"`
	// 			URL  string `json:"url"`
	// 		} `json:"move_learn_method"`
	// 		VersionGroup struct {
	// 			Name string `json:"name"`
	// 			URL  string `json:"url"`
	// 		} `json:"version_group"`
	// 	} `json:"version_group_details"`
	// } `json:"moves"`
	Name string `json:"name"`
	// Order         int    `json:"order"`
	// PastAbilities []any  `json:"past_abilities"`
	// PastTypes     []struct {
	// 	Generation struct {
	// 		Name string `json:"name"`
	// 		URL  string `json:"url"`
	// 	} `json:"generation"`
	// 	Types []struct {
	// 		Slot int `json:"slot"`
	// 		Type struct {
	// 			Name string `json:"name"`
	// 			URL  string `json:"url"`
	// 		} `json:"type"`
	// 	} `json:"types"`
	// } `json:"past_types"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func (c *Client) GetAreas(url string) (areaLocationResponse, error) {
	// sends http request to get area, returns, arealocationresponse struct

	// check if url is in cache
	val, ok := c.cache.Get(url)
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
	c.cache.Add(url, data)
	areas := areaLocationResponse{}
	if err := json.Unmarshal(data, &areas); err != nil {
		return areaLocationResponse{}, errors.New("error while decoding to json data to struct")
	}
	return areas, nil

}
func (c *Client) GetAreaPokemons(area string) (areaLocationDetailedResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + area
	val, ok := c.cache.Get(url)
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
	c.cache.Add(url, data)
	areaData := areaLocationDetailedResponse{}
	err = json.Unmarshal(data, &areaData)
	if err != nil {
		return areaLocationDetailedResponse{}, err
	}
	return areaData, nil

}
func (c *Client) GetPokemon(pokemon string) (PokemonResponse, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	val, ok := c.cache.Get(url)
	if ok {
		pokemonData := PokemonResponse{}
		if err := json.Unmarshal(val, &pokemonData); err != nil {
			return PokemonResponse{}, err
		}
		return pokemonData, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return PokemonResponse{}, errors.New("request unsucessful. check if pokemon specified is correct")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}
	c.cache.Add(url, data)
	pokemonData := PokemonResponse{}
	if err := json.Unmarshal(data, &pokemonData); err != nil {
		return PokemonResponse{}, err
	}
	return pokemonData, nil

}
