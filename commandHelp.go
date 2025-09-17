package main

import (
	"fmt"
)

func commandHelp(config *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range allCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
