package main

import (
	"fmt"
	"errors"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Must specify which pokémon to inspect")
	}
	pokemonName := args[0]
	pokemon, exist := cfg.pokedex[pokemonName]
	if exist == false {
		return errors.New("You have not cought that pokémon yet")
	}
	fmt.Printf(`
Name: %s
Height: %d
Stats:
`, pokemon.Name, pokemon.Height)
	for _, stats := range pokemon.Stats {
		fmt.Printf("  %s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("  - %s\n", types.Type.Name)
	}
	return nil
}