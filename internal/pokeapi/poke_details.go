package pokeapi

import (
	"io"
	"fmt"
	"encoding/json"
)

func (c *Client) GetPokeDetails(pokeName string) (Pokemon, error) {
	url := baseURL + "pokemon/" + pokeName

	// check cache before HTTP call
	if cached, ok := c.cache.Get(url); ok {
		var result Pokemon
		if err := json.Unmarshal(cached, &result); err != nil {
			return Pokemon{}, err
		}
		return result, nil
	}

	// do HTTP call
	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("status code error: %d", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var result Pokemon
	if err := json.Unmarshal(body, &result); err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, body)
	return result, nil
}