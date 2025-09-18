package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, areaName ...string) error {
	if len(areaName) < 1 {
		return errors.New("you must provide a location name!")
	}
	locationSpecificResponse, err := cfg.pokeapiClient.ListLocationSpecific(areaName[0])
	if err != nil {
		fmt.Println("Invalid area name!")
		return err
	}
	fmt.Println("Listing the pokémon in", locationSpecificResponse.Name)
	for _, encounter := range locationSpecificResponse.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}