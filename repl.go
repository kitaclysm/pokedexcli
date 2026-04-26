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

type Config struct {
	Next		*string
	Previous	*string
}

// start the CLI
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{}
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
				err := command.callback(cfg)
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
	callback	func(*Config) error
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