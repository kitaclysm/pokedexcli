package main

import (
	"time"

	"github.com/kitaclysm/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &Config{PokeClient: pokeClient}
	startRepl(cfg)
}