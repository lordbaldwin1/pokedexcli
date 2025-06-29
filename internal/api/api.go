package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedData, ok := c.cache.Get(url)
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
		c.cache.Add(url, data)

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

func (c *Client) ExploreLocation(location *string) (ExploreLocationResponse, error) {
	if location == nil {
		return ExploreLocationResponse{}, errors.New("error: no location input")
	}

	url := baseURL + "/location-area/" + *location

	cacheData, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ExploreLocationResponse{}, errors.New("error: failed to make request")
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return ExploreLocationResponse{}, errors.New("error: failed to execute request")
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return ExploreLocationResponse{}, errors.New("error: failed to read request body")
		}

		var locationRes ExploreLocationResponse
		err = json.Unmarshal(data, &locationRes)
		if err != nil {
			return ExploreLocationResponse{}, errors.New("error: failed to decode JSON data into struct")
		}

		c.cache.Add(url, data)

		return locationRes, nil
	}

	var cacheRes ExploreLocationResponse
	err := json.Unmarshal(cacheData, &cacheRes)
	if err != nil {
		return ExploreLocationResponse{}, errors.New("error: failed to decode data from cache")
	}

	return cacheRes, nil
}

func (c *Client) CatchPokemon(name *string) (Pokemon, error) {
	if name == nil {
		return Pokemon{}, errors.New("error: pokemon name not input")
	}

	url := baseURL + "/pokemon/" + *name

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, errors.New("error: failed to create request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, errors.New("error: request failed")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, errors.New("error: failed to parse response body")
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, errors.New("error: failed to decode response data")
	}

	return pokemon, nil
}
