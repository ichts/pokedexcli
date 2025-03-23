package main

import (
	"fmt"
)

func commandInspect(config *cfg) error {
	pmName := config.Flag

	pm, exists := pokedex[pmName]

	if exists {
		fmt.Println("Name:", pm.Name)
		fmt.Println("Height:", pm.Height)
		fmt.Println("Weight:", pm.Weight)
		fmt.Println("Stats:")

		for _, stat := range pm.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range pm.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
	}

	return nil
}
