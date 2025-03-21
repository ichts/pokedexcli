package main

import (
	"fmt"
	"github.com/ichts/pokedexcli/internal/pokeapi"
)

var pokedex = make(map[string]pokeapi.Pokemon)

func addPokemon(pm pokeapi.Pokemon) {
	pokedex[pm.Name] = pm
	for _, v := range pokedex {
		fmt.Printf("Added to pokedex\nid: %d, name: %s, base_experience: %d\n", v.ID, v.Name, v.BaseExperience)
	}
}
