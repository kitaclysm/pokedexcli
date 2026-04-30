package main

import (
	"fmt"
)

// INSPECT command
func commandInspect(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	pokeName := args[0]
	if info, ok := cfg.Pokedex[pokeName]; ok {
		fmt.Printf("Name: %s\n", info.Name)
		fmt.Printf("Height: %d\n", info.Height)
		fmt.Printf("Weight: %d\n", info.Weight)
		fmt.Println("Stats:")
		for _, stat := range info.Stats {
			fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pType := range info.Types {
			fmt.Printf(" - %s\n", pType.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}