package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lordbaldwin1/pokedexcli/internal/api"
)

type config struct {
	pokeapiClient api.Client
	nextURL       *string
	prevURL       *string
}

func startRepl(cfg *config) {
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
			command.callback(cfg)
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
