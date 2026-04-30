package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"

	"github.com/kitaclysm/pokedexcli/internal/pokeapi"
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

type Config struct {
	PokeClient	pokeapi.Client
	Next		*string
	Previous	*string
	Pokedex		map[string]pokeapi.Pokemon
}

// start the CLI
func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	// cfg := &Config{}
	for ;; {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			results := cleanInput(userInput)
			if len(results) == 0 {
				continue
			}
			userCall := results[0]
			args := results[1:]
			command, exists := getCommands()[userCall]
			if exists {
				err := command.callback(cfg, args...)
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
	callback	func(*Config, ...string) error
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
		"explore": {
			name:			"explore",
			description: 	"Displays all pokemon in specific area (use: explore <area-name>)",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Attempts to catch a specified Pokemon. On success, Pokemon is added to Pokedex",
			callback:		commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description:	"Displays information for pokemon that have been caught",
			callback:		commandInspect,
		},
	}
}