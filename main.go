package main

import (
	"time"

	"github.com/lordbaldwin1/pokedexcli/internal/api"
)

func main() {
	pokeClient := api.NewClient(5*time.Second, 30*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
