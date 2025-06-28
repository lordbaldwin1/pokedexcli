package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := config{
		nextUrl: "https://pokeapi.co/api/v2/location-area/?offset=0",
		prevUrl: "",
	}

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
			command.callback(&cfg)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	strSlice := strings.Split(strings.ToLower(text), " ")

	var res []string

	for _, str := range strSlice {
		trimmed := strings.TrimSpace(str)

		if trimmed != "" {
			res = append(res, trimmed)
		}
	}

	return res
}

type config struct {
	nextUrl string
	prevUrl string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
