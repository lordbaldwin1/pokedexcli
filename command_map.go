package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {
	res, err := http.Get(cfg.nextUrl)
	if err != nil {
		return errors.New("error: request to pokeapi failed")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("error: failed to read response into memory")
	}

	var locations LocationAreaResponse
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return errors.New("error: failed to decode response JSON bytes")
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.nextUrl = locations.Next
	if locations.Previous != nil {
		cfg.prevUrl = *locations.Previous
	} else {
		cfg.prevUrl = ""
	}

	return nil
}
