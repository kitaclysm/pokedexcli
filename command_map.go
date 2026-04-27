package main

import (
	"fmt"
)

// MAP command
func commandMap(cfg *Config) error {
	res, err := cfg.PokeClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = res.Next
	cfg.Previous = res.Previous
	for _, item := range res.Results {
		fmt.Println(item.Name)
	}
	return nil
}

// MAPB command
func commandMapB(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := cfg.PokeClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Previous = res.Previous
	cfg.Next = res.Next
	for _, item := range res.Results {
		fmt.Println(item.Name)
	}
	return nil
}