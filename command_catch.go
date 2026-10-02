package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid command")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	baseExperience := pokemon.BaseExperience
	luck := rand.IntN(1000)

	if luck > baseExperience {
		fmt.Printf("%s was caught!\n", name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
		return nil
	}

	fmt.Printf("%s escaped!\n", name)

	return nil
}
