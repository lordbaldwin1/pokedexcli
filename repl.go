package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lordbaldwin1/pokedexcli/internal/api"
	"github.com/lordbaldwin1/pokedexcli/internal/cache"
)

type config struct {
	pokeapiClient api.Client
	nextURL       *string
	prevURL       *string
}

func startRepl(cfg *config, cache *cache.Cache) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Poxedex > ")
		scanner.Scan()

		cleanedLine := cleanInput(scanner.Text())
		if len(cleanedLine) == 0 {
			continue
		}
		commandName := cleanedLine[0]

		commandRegistry := getCommands()

		command, exists := commandRegistry[commandName]
		if exists {
			command.callback(cfg, cache)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.ToLower(text)
	cleanedText := strings.Fields(trimmed)
	return cleanedText
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *cache.Cache) error
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
	}
}
