package main

import (
	"fmt"
)

// HELP callback
func commandHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range cfg.Commands {
		fmt.Printf("%s:	%s\n", cmd.name, cmd.description)
	}
	return nil
}