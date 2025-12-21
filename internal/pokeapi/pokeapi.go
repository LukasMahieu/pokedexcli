package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLocations(url string) (LocationAPI, error) {
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

	var location LocationAPI
	if err := json.Unmarshal(body, &location); err != nil {
		return LocationAPI{}, err
	}
	return location, nil
}
