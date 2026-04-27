package pokeapi

import (
	"io"
	"fmt"
	"encoding/json"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL
	if pageURL != nil {
		url = *pageURL
	}

	// check cache before HTTP call
	if cached, ok := c.cache.Get(url); ok {
		var result RespShallowLocations
		if err := json.Unmarshal(cached, &result); err != nil {
			return RespShallowLocations{}, err
		}
		return result, nil
	}

	// do HTTP call
	res, err := c.httpClient.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	if res.StatusCode > 299 {
		return RespShallowLocations{}, fmt.Errorf("status code error: %d", res.StatusCode)
	}

	var result RespShallowLocations
	if err := json.Unmarshal(body, &result); err != nil {
		return RespShallowLocations{}, err
	}
	c.cache.Add(url, body)
	return result, nil
}