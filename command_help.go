package main

import "fmt"

func commandHelp(config *cfg) error {
	fmt.Println("Welcome to the Pokedex!")

	fmt.Println("Usage:")
	fmt.Println()

	for _, nd := range namesAndDescriptions {
		fmt.Printf("%s: %s\n", nd.name, nd.description)
	}

	return nil
}
