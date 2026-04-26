package main

import (
	"fmt"
)

// HELP callback
func commandHelp(cnfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s:	%s\n", cmd.name, cmd.description)
	}
	return nil
}