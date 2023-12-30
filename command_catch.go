package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemon := args[0]
	fmt.Printf("Exploring %v...\n", pokemon)
	pokemonResp, err := c.pokeapiClient.GetPokemon(&pokemon)

	if err != nil {
		return err
	}

	chance := rand.Intn(pokemonResp.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)

	if chance > 40 {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonResp.Name)

	c.caughtPokemon[pokemonResp.Name] = pokemonResp
	return nil
}
