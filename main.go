package main

import (
	"net/http"
	"time"

	"github.com/ichts/pokedexcli/internal/pokeapi"
	"github.com/ichts/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*cfg) error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays next 20 locations",
		callback:    commandMapf,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays previous 20 locations",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "List all the Pokemon located in a location area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Catch a pokemon",
		callback:    commandCatch,
	},
}

type nameAndDescription struct {
	name        string
	description string
}

type cfg struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	Cache         *pokecache.Cache
	Flag          string
}

var namesAndDescriptions = []nameAndDescription{
	{
		name:        "exit",
		description: "Exit the Pokedex",
	},
	{
		name:        "help",
		description: "Displays a help message",
	},
	{
		name:        "map",
		description: "Displays next 20 locations",
	},
}

func main() {
	pokeClient := pokeapi.Client{
		HttpClient: http.Client{},
	}

	itv := time.Second * 10
	c := pokecache.NewCache(itv)

	config := &cfg{
		pokeapiClient: pokeClient,
		Cache:         c,
	}
	startRepl(config)
}
