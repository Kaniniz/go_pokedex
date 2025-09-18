package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Must specify which pokémon to catch!")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if rand.Intn(pokemon.BaseExperience) < pokemon.BaseExperience/2 {
		fmt.Printf("The %s ran away!\n", pokemon.Name)
		return err
	}
	

	fmt.Printf("%s was cought!\n", pokemon.Name)
	cfg.pokedex[pokemon.Name] = pokemon
	return nil
}