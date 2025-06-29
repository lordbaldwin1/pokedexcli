package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/lordbaldwin1/pokedexcli/internal/cache"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string, cache *cache.Cache) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedData, ok := cache.Get(url)
	if !ok {
		// create request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaResponse{}, errors.New("error: failed to create request")
		}

		// execute request
		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResponse{}, errors.New("error: failed to execute request")
		}
		defer res.Body.Close()

		// read all bytes of data from body
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaResponse{}, errors.New("error: failed to read data from response")
		}

		// cache byte data
		cache.Add(url, data)

		// decode data and unmarshal it into struct
		var locations LocationAreaResponse
		err = json.Unmarshal(data, &locations)
		if err != nil {
			return LocationAreaResponse{}, errors.New("error: failed to decode data")
		}
		return locations, nil
	}

	var cachedLocations LocationAreaResponse
	err := json.Unmarshal(cachedData, &cachedLocations)
	if err != nil {
		return LocationAreaResponse{}, errors.New("error: cached data failed to decode")
	}

	return cachedLocations, nil
}
