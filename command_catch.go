package main

import (
	"fmt"
	"math/rand"
	"math"
)

func commandCatchPokemon(cfg *config) error {
	
	pokeResp, err := cfg.pokeapiClient.GetPokemonData(cfg.commandParameter)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeResp.Name)
	pokeExp := pokeResp.BaseExperience

	catchProbability := calculateCatchChanceExponential(pokeExp)

	roll := rand.Float64()

	if roll < catchProbability {
		fmt.Printf("%s was caught!\n", pokeResp.Name)
		cfg.caughtPokemon[pokeResp.Name] = pokeResp
	} else {
		fmt.Printf("%s escaped!\n", pokeResp.Name)
	}

	return nil
}

func calculateCatchChanceExponential(xp int) float64 {
	const decayRate = 0.01 

	catchChance := math.Exp(-decayRate * float64(xp))

	return catchChance
}
