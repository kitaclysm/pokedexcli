package main

import (
	"fmt"
	"os"
)

// EXIT callback
func commandExit(cnfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}