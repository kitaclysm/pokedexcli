package pokeapi

// JSON result location-areas -> list areas
type RespShallowLocations struct {
	Count    int    	`json:"count"`
	Next     *string	`json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string 	`json:"name"`
		URL  string 	`json:"url"`
	} `json:"results"`
}

// JSON result specific area -> list pokemon found in area
type LocationEncounters struct {
	Name  string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}