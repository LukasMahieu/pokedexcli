package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/LukasMahieu/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

func FetchLocations(url string, cache pokecache.Cache) (LocationAPI, error) {
	if data, ok := cache.Get(url); ok {
		var location LocationAPI
		json.Unmarshal(data, &location)
		return location, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationAPI{}, err
	}
	status := res.StatusCode
	if status != http.StatusOK {
		return LocationAPI{}, fmt.Errorf("Unexpected status %d", status)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAPI{}, err
	}

	cache.Add(url, body)

	var location LocationAPI
	if err := json.Unmarshal(body, &location); err != nil {
		return LocationAPI{}, err
	}
	return location, nil
}
func FetchSpecificLocation(url string, cache pokecache.Cache) (SpecificLocationAPI, error) {
	if data, ok := cache.Get(url); ok {
		var location SpecificLocationAPI
		json.Unmarshal(data, &location)
		return location, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return SpecificLocationAPI{}, err
	}
	status := res.StatusCode
	if status != http.StatusOK {
		return SpecificLocationAPI{}, fmt.Errorf("Unexpected status %d", status)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return SpecificLocationAPI{}, err
	}

	cache.Add(url, body)

	var location SpecificLocationAPI
	if err := json.Unmarshal(body, &location); err != nil {
		return SpecificLocationAPI{}, err
	}
	return location, nil
}
