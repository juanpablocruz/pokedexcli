package main

import (
	"fmt"
)

func commandPokedex(c *config, args ...string) error {

	fmt.Println("Welcome to the Pokedex!")
	if c.caughtPokemon == nil {
		fmt.Println("Your Pokedex is empty.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
