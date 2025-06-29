package main

import (
	"fmt"
	"os"

	"github.com/lordbaldwin1/pokedexcli/internal/cache"
)

func commandHelp(cfg *config, cache *cache.Cache) error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	fmt.Print("\n")

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}

func commandExit(cfg *config, cache *cache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, cache *cache.Cache) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.nextURL, cache)
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

func commandMapB(cfg *config, cache *cache.Cache) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.prevURL, cache)
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
