package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range allCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}