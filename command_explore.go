package main

import (
	"fmt"
)

func commandExplore(config *cfg) error {
	pokemonList, err := config.pokeapiClient.GetPokemonList(config.Flag, config.Cache)

	fmt.Printf("Exploring %s\n", config.Flag)

	for _, p := range pokemonList {
		fmt.Println(p)
	}

	if err != nil {
		return err
	}

	return nil
}
