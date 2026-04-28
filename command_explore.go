package main

import (
	"fmt"
)

// EXPLORE command
func commandExplore(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location area name")
	}
	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)
	res, err := cfg.PokeClient.GetAreaPokemon(areaName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, creature := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", creature.Pokemon.Name)
	}
	return nil
}