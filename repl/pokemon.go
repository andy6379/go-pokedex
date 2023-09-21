package repl

import (
	"awesomeProject4/internal/pokeapi"
	"awesomeProject4/model"
	"fmt"
)

var caughtPokemon = make(map[string]model.Pokemon)

func catchPokemon(pokemonName string) error {

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// Fetch Pokemon data from the PokeAPI
	pokemon, err := pokeapi.GetPokemonByName(pokemonName) // This function needs to be implemented
	if err != nil {
		return err
	}

	// Calculate the chance of catching the Pokemon based on its base experience
	catchChance := 100.0 / float64(1+pokemon.BaseExperience)

	// Generate a random number between 0 and 100
	//randChance := rand.Float64() * 100.0
	randChance := 0.01

	if randChance <= catchChance {
		caughtPokemon[pokemon.Name] = *pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func inspect(pokemonName string) error {
	pokemon, exists := caughtPokemon[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	// Display the details of the caught Pokémon
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func pokedex() error {
	if len(caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
