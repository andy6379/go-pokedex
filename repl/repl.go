package repl

import (
	"awesomeProject4/model"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl(cfg *model.Config) {
	commands := getCommands(cfg)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputParts := strings.Fields(input) // Split by whitespace
		if len(inputParts) == 0 {
			continue
		}

		commandName := inputParts[0]
		param := ""
		if command, exists := commands[commandName]; exists {
			if command.Callback != nil && len(inputParts) > 1 {
				// Run the command with the parameter
				if err := command.Callback(strings.Join(inputParts[1:], " ")); err != nil {
					fmt.Printf("Error executing command: %v\n", err)
				}
			} else if command.Callback != nil {
				// Run the command without a parameter
				if err := command.Callback(param); err != nil {
					fmt.Printf("Error executing command: %v\n", err)
				}
			} else {
				fmt.Println("Invalid command usage. Type 'help' for guidance.")
			}
		} else {
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}

func getCommands(config *model.Config) map[string]model.CliCommand {
	return map[string]model.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    func(_ string) error { return commandHelp(config) },
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    func(_ string) error { return commandExit() },
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 location areas",
			Callback:    func(_ string) error { return commandMap(config) },
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 location areas",
			Callback:    func(_ string) error { return commandMapBack(config) },
		},
		"explore": {
			Name:        "explore <location_name>",
			Description: "Explore a specific Pokemon area",
			Callback:    exploreArea,
		},
		"catch": {
			Name:        "catch <pokemon_name>",
			Description: "Attempt to catch a pokemon",
			Callback:    catchPokemon,
		},
		"inspect": {
			Name:        "inspect <pokemon_name>",
			Description: "Inspect a caught pokemon",
			Callback:    inspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Displays all caught pokemon",
			Callback:    func(_ string) error { return pokedex() },
		},
	}
}
