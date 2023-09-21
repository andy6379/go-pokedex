package model

type PokemonEncounters struct {
	Pokemon PokemonInstance `json:"pokemon"`
}

type PokemonInstance struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationDetailResponse struct {
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}

type baseDetail struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Stat struct {
	Stat     baseDetail `json:"stat"`
	BaseStat int        `json:"base_stat"`
}

type Type struct {
	Type baseDetail `json:"type"`
}
