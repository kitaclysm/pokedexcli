package pokeapi

import (
	"net/http"
	"io"
	"fmt"
	"encoding/json"
)

// JSON result location-areas
type Location struct {
	Count    int    	`json:"count"`
	Next     *string	`json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string 	`json:"name"`
		URL  string 	`json:"url"`
	} `json:"results"`
}

func GetLocations(pageURL *string) (Location, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != nil {
		url = *pageURL
	}
	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}
	if res.StatusCode > 299 {
		return Location{}, fmt.Errorf("status code error: %d", res.StatusCode)
	}
	var result Location
	if err := json.Unmarshal(body, &result); err != nil {
		return Location{}, err
	}

	return result, nil
}
