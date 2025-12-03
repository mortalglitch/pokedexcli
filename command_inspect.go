package main

import (
	"fmt"
)

func commandInspect(cfg *config) error {
	searchParameter := *cfg.commandParameter
	val, ok := cfg.caughtPokemon[searchParameter]
	if ok {
		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Weight: %d\n", val.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range val.Stats{
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, types := range val.Types {
			fmt.Printf("  - %s\n", types.Type.Name)
		}
	} else {
		fmt.Println("No Pokemon by that name has been caught!")
	}
	return nil
}


