package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/lordbaldwin1/pokedexcli/internal/api"
)

type cliCommand struct {
	name        string
	description string
	usage       string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			usage:       "exit",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			usage:       "help",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 areas in the Pokemon world.",
			usage:       "map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 areas in the Pokemon world.",
			usage:       "mapb",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Show the Pokemon in a specific location",
			usage:       "explore <area-name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			usage:       "catch <pokemon_name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a currently caught Pokemon",
			usage:       "inspect <pokemon_name>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View your Pokedex",
			usage:       "pokedex, OR pokedex <pokemon_name>",
			callback:    commandPokedex,
		},
	}
}

func commandHelp(cfg *config, parameter *string) error {
	fmt.Print("\nWelcome to the Pokedex!\n\nCommands:\n")
	fmt.Print("\n")

	for _, command := range getCommands() {
		fmt.Printf("%s:\n", command.name)
		fmt.Printf("   %s\n", command.description)
		fmt.Printf(" - Usage: %s\n", command.usage)
	}
	fmt.Println("")
	return nil
}

func commandExit(cfg *config, parameter *string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, parameter *string) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)
	if err != nil {
		fmt.Println(err)
	}

	cfg.nextURL = locationRes.Next
	cfg.prevURL = locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(cfg *config, parameter *string) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.prevURL)
	if err != nil {
		fmt.Println(err)
	}

	cfg.nextURL = locationRes.Next
	cfg.prevURL = locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(cfg *config, location *string) error {
	if location == nil {
		fmt.Println("Usage: explore <area_name>")
		return nil
	}
	locationRes, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		fmt.Println(err)
	}

	for _, encounter := range locationRes.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, parameter *string) error {
	if parameter == nil {
		fmt.Println("Usage: explore <pokemon_name")
		return nil
	}

	pokemonRes, err := cfg.pokeapiClient.CatchPokemon(parameter)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	catchChance := calcCatchChance(pokemonRes)
	fmt.Printf("Throwing a Pokeball at %s...\n", *parameter)
	randomNumber := rand.Intn(100)

	var caught bool
	if randomNumber < catchChance {
		caught = true
	} else {
		caught = false
	}

	if caught {
		cfg.pokedex[*parameter] = pokemonRes
		fmt.Printf("%s was caught!\n", *parameter)
		return nil
	}

	fmt.Printf("%s escaped!\n", *parameter)
	return nil
}

func commandInspect(cfg *config, pokemon *string) error {
	if pokemon == nil {
		fmt.Println("Usage: inspect <pokemon_name>")
		return nil
	}

	pokemonData, ok := cfg.pokedex[*pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)

	fmt.Println("Stats:")
	fmt.Printf("	-hp: %d\n", pokemonData.Stats[0].BaseStat)
	fmt.Printf("	-attack: %d\n", pokemonData.Stats[1].BaseStat)
	fmt.Printf("	-defense: %d\n", pokemonData.Stats[2].BaseStat)
	fmt.Printf("	-special-attack: %d\n", pokemonData.Stats[3].BaseStat)
	fmt.Printf("	-special-defense: %d\n", pokemonData.Stats[4].BaseStat)
	fmt.Printf("	-speed: %d\n", pokemonData.Stats[5].BaseStat)

	fmt.Println("Types:")
	for _, pType := range pokemonData.Types {
		fmt.Printf("	-%s\n", pType.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, parameter *string) error {
	if parameter != nil {
		_, ok := cfg.pokedex[*parameter]
		if !ok {
			fmt.Printf("You have not caught %s!\n", *parameter)
			return nil
		}
		fmt.Printf("You have caught %s!\n", *parameter)
	}

	fmt.Println("Your Pokedex:")
	for k := range cfg.pokedex {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}

func calcCatchChance(p api.Pokemon) int {
	const maxXP = 300
	const baseCatchRate = 95
	const minCatchRate = 50

	clampedXP := max(min(p.BaseExperience, maxXP), 1)

	difficultyFactor := float64(clampedXP) / float64(maxXP)
	catchChance := min(max(baseCatchRate-int(difficultyFactor*float64(baseCatchRate-minCatchRate)), minCatchRate), baseCatchRate)

	return catchChance
}
