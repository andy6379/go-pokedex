package repl

import (
	"awesomeProject4/internal/pokeapi"
	"awesomeProject4/model"
	"fmt"
	"os"
)

func commandHelp(cfg *model.Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands(cfg) {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *model.Config) error {
	response, err := pokeapi.GetLocationAreas(config.NextURL)
	if err != nil {
		return err
	}

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	if response.Next != nil {
		config.NextURL = *response.Next
	}

	if response.Previous != nil {
		config.PreviousURL = *response.Previous
	}

	return nil
}

func commandMapBack(config *model.Config) error {
	if config.PreviousURL == "" {
		fmt.Println("You're on the first page!")
		return nil
	}

	response, err := pokeapi.GetLocationAreas(config.PreviousURL)
	if err != nil {
		return err
	}

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	if response.Next != nil {
		config.NextURL = *response.Next
	}

	if response.Previous != nil {
		config.PreviousURL = *response.Previous
	}

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func exploreArea(areaName string) error {
	fmt.Printf("Exploring %s...\n", areaName)

	response, err := pokeapi.GetLocationDetails(areaName) // A new function to fetch details of a specific location
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon: ")
	for _, encounter := range response.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
