# pokedexcli

![Image of welcome message in terminal ascii art text](images/your-image-name.png)
This is a small CLI to allow users to explore the world of Pokemon, find which Pokemon can be encountered in an area, catch Pokemon, inspect the stats of Pokemon, and create a Pokedex out of caught Pokemon!

## Goals

The primary goal for this project was to learn the http standard library, JSON parsing, and caching API responses in Go.

## Reflection

After completing this project, I am now comfortable with the http standard library, JSON parsing, concurrency, and caching in Go. Specifically, I understand how to create an http client, create requests, read request bodies into memory, and unmarshal []byte data into Go Structs.

For caching, I learned how to cache API responses with a map struct, where the key is the URL and the value is a struct holding the creation time and the byte data. This allowed me to cache responses of API requests regardless of the JSON form of the data. When a request would be made, the cache is checked and if the data exists within the cache, that data is returned. This greatly reduces the time to fetch data on repeated calls. If the data is cached, the byte data from the response is stored in the cache. A lifetime was added to the cache and records are reaped after that interval has passed. This checking of the cache runs concurrently with all other CLI behavior.

This was my first project in Go, so I've also learned how to structure a project in Go using internal packages to separate things like the API interactions and caching behavior. I've also learned how to practically understand and handle the zero values for various data types, including maps.

## Testing

Basic unit tests are added to test crucial behavior for input parsing and cache behavior. These tests can be found in `repl_test.go` and `/internal/cache/cache_test.go`. Tests were written using Go's standard testing library.
