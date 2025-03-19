package main

import (
	"fmt"
)

func commandMapf(config *cfg) error {
	locationList, err := config.pokeapiClient.GetLocationList(config.Next, config.Cache)
	if err != nil {
		return err
	}

	config.Next = &locationList.Next
	config.Previous = &locationList.Previous

	return nil
}

func commandMapb(config *cfg) error {
	if *config.Previous == "" {
		fmt.Println("you are on the first page")
		return nil
	}
	locationList, err := config.pokeapiClient.GetLocationList(config.Previous, config.Cache)
	if err != nil {
		return err
	}

	config.Next = &locationList.Next
	config.Previous = &locationList.Previous

	return nil
}
