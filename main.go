package main

import (
	"awesomeProject4/internal/pokeapi"
	"awesomeProject4/model"
	"awesomeProject4/repl"
)

func main() {
	config := &model.Config{NextURL: pokeapi.BaseURL}
	repl.StartRepl(config)
}
