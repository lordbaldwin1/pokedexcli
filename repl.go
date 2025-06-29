package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lordbaldwin1/pokedexcli/internal/api"
)

const asciiString = `
		⢰⣶⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀
		⠀⣿⣿⣿⣷⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣶⣾⣿
		⠀⠘⢿⣿⣿⣿⣿⣦⣀⣀⣀⣄⣀⣀⣠⣀⣤⣶⣿⣿⣿⣿⣿⠇
		⠀⠀⠈⠻⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠀
		⠀⠀⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⠋⠀⠀⠀
		⠀⠀⠀⢠⣿⣿⡏⠆⢹⣿⣿⣿⣿⣿⣿⠒⠈⣿⣿⣿⣇⠀⠀⠀
		⠀⠀⠀⣼⣿⣿⣷⣶⣿⣿⣛⣻⣿⣿⣿⣶⣾⣿⣿⣿⣿⡀⠀⠀
		⠀⠀⠀⡁⠀⠈⣿⣿⣿⣿⢟⣛⡻⣿⣿⣿⣟⠀⠀⠈⣿⡇⠀⠀
		⠀⠀⠀⢿⣶⣿⣿⣿⣿⣿⡻⣿⡿⣿⣿⣿⣿⣶⣶⣾⣿⣿⠀⠀
		⠀⠀⠀⠘⣿⣿⣿⣿⣿⣿⣿⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠀
		⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀
  ____   ___  _  _______ ____  _______  ______ _     ___ 
 |  _ \ / _ \| |/ / ____|  _ \| ____\ \/ / ___| |   |_ _|
 | |_) | | | | ' /|  _| | | | |  _|  \  / |   | |    | | 
 |  __/| |_| | . \| |___| |_| | |___ /  \ |___| |___ | | 
 |_|    \___/|_|\_\_____|____/|_____/_/\_\____|_____|___|

`

type config struct {
	pokeapiClient api.Client
	nextURL       *string
	prevURL       *string
	pokedex       map[string]api.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(asciiString)
	fmt.Println("Type 'help' to see a list of commands")
	fmt.Println()
	for {
		fmt.Print("Poxedex > ")
		scanner.Scan()

		cleanedLine := cleanInput(scanner.Text())
		if len(cleanedLine) == 0 {
			continue
		}
		commandName := cleanedLine[0]
		var commandArg *string
		if len(cleanedLine) > 1 {
			commandArg = &cleanedLine[1]
		} else {
			commandArg = nil
		}

		commandRegistry := getCommands()

		command, exists := commandRegistry[commandName]
		if exists {
			command.callback(cfg, commandArg)
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
