package main

import (
	"errors"
	"fmt"
)

func commandMapf (cfg *config) error {
	locationsAreaResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsAreaResponse.Next
	cfg.prevLocationsURL = locationsAreaResponse.Previous

	for _, loc := range locationsAreaResponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
} 

func commandMapb (cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsAreaResponse, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsAreaResponse.Next
	cfg.prevLocationsURL = locationsAreaResponse.Previous

	for _, loc := range locationsAreaResponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}