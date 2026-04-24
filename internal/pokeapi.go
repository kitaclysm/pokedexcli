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
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Println("error getting data")
		return Location{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading data")
		return Location{}, err
	}
	res.Body.Close()
	if res.StatusCode > 299 {
		return Location{}, fmt.Errorf("status code error: %d", res.StatusCode)
	}
	var result Location
	if err := json.Unmarshal(body, &result); err != nil {
		return Location{}, err
	}

	// Handle this in commandMap and commandMapB functions?
	// for _, entry := range result.Results {
	// 	fmt.Println(entry.Name)
	// }

	return result, nil
}
