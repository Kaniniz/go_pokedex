package main

import (
	"fmt"
)

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range allCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}