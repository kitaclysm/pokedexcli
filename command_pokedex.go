package main

import (
	"fmt"
)

// POKEDEX command
func commandPokedex(cfg *Config, args ...string) error {
	dex := cfg.Pokedex
	if len(dex) > 0 {
		fmt.Println("Your Pokedex:")
		for _, entry := range dex {
			fmt.Printf(" - %s\n", entry.Name)
		}
	} else {
		fmt.Println("you have not caught any pokemon")
	}
	return nil
}