package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 areas in the Pokemon world.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Show the Pokemon at a specific location",
			callback:    commandExplore,
		},
	}
}

func commandHelp(cfg *config) error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	fmt.Print("\n")

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {
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

func commandMapB(cfg *config) error {
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

func commandExplore(cfg *config) error {

	return nil
}
