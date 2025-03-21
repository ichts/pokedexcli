package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(config *cfg) error {
	pm := config.pokeapiClient.GetPokemon(config.Flag)

	fmt.Printf("Throwing a Pokeball at %s...\n", pm.Name)
	dice := rand.IntN(250)
	caught := (dice >= pm.BaseExperience)
	if caught {
		fmt.Printf("%s was caught!\n", config.Flag)
		addPokemon(pm)
		return nil
	}

	fmt.Printf("%s escaped!\n", config.Flag)
	return nil
}
