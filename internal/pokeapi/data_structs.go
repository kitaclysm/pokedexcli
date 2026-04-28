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

// JSON results specific pokemon data
type Pokemon struct {
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Name string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}