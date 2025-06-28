package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(cfg *config) error {
	res, err := http.Get(cfg.prevUrl)
	if err != nil {
		return errors.New("error: failed to get response from pokeapi")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("error: failed to read data from response")
	}

	var locations LocationAreaResponse
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return errors.New("error: failed to decode JSON bytes")
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
