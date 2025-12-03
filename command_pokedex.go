package main

import (
	"fmt"
)

func commandPokedex(cfg *config) error {
	fmt.Println("Your Pokedex:")
	for _, poke := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", poke.Name)
	}
	return nil
}	
