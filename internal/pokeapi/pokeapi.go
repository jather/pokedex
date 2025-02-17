package pokeapi

import (
	"encoding/json"
	"errors"
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

func (c *Client) GetAreas(url string) (areaLocationResponse, error) {
	// sends http request to get area, returns, arealocationresponse struct
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return areaLocationResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return areaLocationResponse{}, nil
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	areas := areaLocationResponse{}
	if err := decoder.Decode(&areas); err != nil {
		return areaLocationResponse{}, errors.New("error while decoding to json data to struct")
	}
	return areas, nil

}
