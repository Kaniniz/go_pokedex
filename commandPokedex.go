package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {
	fmt.Println("Your pokédex entries:")
	for _, entry := range cfg.pokedex {
		fmt.Println(entry.Name)
	}
	return nil
}