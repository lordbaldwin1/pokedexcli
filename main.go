package main

import (
	"time"

	"github.com/lordbaldwin1/pokedexcli/internal/api"
	"github.com/lordbaldwin1/pokedexcli/internal/cache"
)

func main() {
	pokeClient := api.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	cache := cache.NewCache(60 * time.Second)
	startRepl(cfg, cache)
}
