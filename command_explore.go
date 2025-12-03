package main

import (
	"fmt"
)

func commandExploreArea(cfg *config) error {
	
	areaResp, err := cfg.pokeapiClient.ListAreaPokemon(cfg.commandParameter)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s... \nFound Pokemon: \n", areaResp.Name)
	for _, poke := range areaResp.PokemonEncounters {
		fmt.Println(poke.Pokemon.Name)
	}
	return nil
}


