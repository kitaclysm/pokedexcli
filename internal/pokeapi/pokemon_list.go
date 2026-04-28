package pokeapi

import (
	"io"
	"fmt"
	"encoding/json"
)

func (c *Client) GetAreaPokemon(areaName string) (LocationEncounters, error) {
	url := baseURL + "location-area/" + areaName

	// check cache before HTTP call
	if cached, ok := c.cache.Get(url); ok {
		var result LocationEncounters
		if err := json.Unmarshal(cached, &result); err != nil {
			return LocationEncounters{}, err
		}
		return result, nil
	}

	// do HTTP call
	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationEncounters{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationEncounters{}, fmt.Errorf("status code error: %d", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationEncounters{}, err
	}

	var result LocationEncounters
	if err := json.Unmarshal(body, &result); err != nil {
		return LocationEncounters{}, err
	}
	c.cache.Add(url, body)
	return result, nil
}