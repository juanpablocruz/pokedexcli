package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	location := args[0]
	fmt.Printf("Exploring %v...\n", location)
	locationResp, err := c.pokeapiClient.GetLocation(&location)

	if err != nil {
		return err
	}

	if len(locationResp.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		for _, pokemonEncounter := range locationResp.PokemonEncounters {
			fmt.Printf("- %s\n", pokemonEncounter.Pokemon.Name)
		}
	}

	return nil
}
