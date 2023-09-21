package model

type CliCommand struct {
	Name        string
	Description string
	Callback    func(param string) error
}

type Config struct {
	NextURL     string
	PreviousURL string
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}
