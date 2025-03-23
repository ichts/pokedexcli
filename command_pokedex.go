package main

import (
	"fmt"
)

func commandPokedex(config *cfg) error {
	fmt.Println("Your Pokedex:")
	for name, _ := range pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
