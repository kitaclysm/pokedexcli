package main

import (
	"fmt"
	"math/rand"
)

// CATCH command
func commandCatch(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	pokeName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)
	res, err := cfg.PokeClient.GetPokeDetails(pokeName)
	if err != nil {
		return err
	}
	baseExp := res.BaseExperience
	threshold := rand.Intn(450)
	if threshold > baseExp {
		cfg.Pokedex[pokeName] = res
		fmt.Printf("%s was caught!\n", pokeName)
	} else {
		fmt.Printf("%s escaped!\n", pokeName)
	}
	return nil
}