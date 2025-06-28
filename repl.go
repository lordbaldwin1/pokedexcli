package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
			command.callback()
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
