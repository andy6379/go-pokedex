package pokeapi

import (
	"awesomeProject4/internal/pokecache"
	"awesomeProject4/model"
	"encoding/json"
	"net/http"
	"time"
)

const (
	BaseURL        = "https://pokeapi.co/api/v2/location-area/"
	PokemonBaseURL = "https://pokeapi.co/api/v2/pokemon/"
)

var apiCache = pokecache.NewCache(5 * time.Minute)

func GetLocationAreas(url string) (*model.LocationAreaResponse, error) {
	if url == "" {
		url = BaseURL
	}

	// Try to get from the cache
	cachedData, exists := apiCache.Get(url)
	if exists {
		var response model.LocationAreaResponse
		if err := json.Unmarshal(cachedData, &response); err != nil {
			return nil, err
		}
		return &response, nil
	}

	// If not in the cache, fetch from the API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response model.LocationAreaResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	// Convert the struct to bytes and add to cache
	responseData, err := json.Marshal(&response)
	if err != nil {
		return nil, err
	}
	apiCache.Add(url, responseData)

	return &response, nil
}

func GetLocationDetails(areaName string) (*model.LocationDetailResponse, error) {
	url := BaseURL + areaName
	cachedData, exists := apiCache.Get(url)
	if exists {
		var response model.LocationDetailResponse
		if err := json.Unmarshal(cachedData, &response); err != nil {
			return nil, err
		}
		return &response, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response model.LocationDetailResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	responseData, err := json.Marshal(&response)
	if err != nil {
		return nil, err
	}
	apiCache.Add(url, responseData)

	return &response, nil
}

func GetPokemonByName(name string) (*model.Pokemon, error) {
	url := PokemonBaseURL + name

	cachedData, exists := apiCache.Get(url)
	if exists {
		var pokemon model.Pokemon
		if err := json.Unmarshal(cachedData, &pokemon); err != nil {
			return nil, err
		}
		return &pokemon, nil
	}

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemon model.Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, err
	}

	return &pokemon, nil
}
