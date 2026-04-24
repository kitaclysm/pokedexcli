package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

// standardize input
func cleanInput(text string) []string {
	var words []string
	fields := strings.Fields(text)
	if len(fields) == 0 {
		return []string{}
	}
	for _, field := range fields {
		field = strings.ToLower(field)
		words = append(words, field)
	}
	return words
}

// start the CLI
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			results := cleanInput(userInput)
			if len(results) == 0 {
				continue
			}
			userCall := results[0]
			command, exists := getCommands()[userCall]
			if exists {
				err := command.callback()
				if err != nil {
					fmt.Println(err)
				}

			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

// commands registry structure
type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

// defined commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"map": {
			name:			"map",
			description:	"Displays next 20 location areas in Pokemon world",
			callback: 		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Displays previous 20 location areas in Pokemon world",
			callback:		commandMapB,
		},
	}
}

type Config struct {
	Next		*string
	Previous	*string
}

// EXIT callback
func commandExit(cnfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

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

// MAP command
func commandMap(cnfg *Config) error {
	var result Location
	if cnfg.Next == nil {
		result = GetLocations("https://pokeapi.co/api/v2/location-area/")
		cnfg.Next = result.Next
		cnfg.Previous = result.Previous
	} else {
		url = *pageURL
	}
}

// MAPB command
func commandMapB(cnfg *Config) error {

}