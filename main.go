package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Kaniniz/go_Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	response := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		response.Scan()
		cleanText := cleanInput(response.Text())
		if err := response.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		//First word in the user input is the command
		input := cleanText[0]
		command, exists := allCommands[input]
		if exists == false {
			fmt.Println("Unknown command:", input)
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}

	}
}

func cleanInput(text string) []string {
	text, _ = strings.CutPrefix(text, " ")
	text, _ = strings.CutSuffix(text, " ")
	text = strings.ToLower(text)
	words := strings.Split(text, " ")
	return words
}

type cliCommands struct {
	name        string
	description string
	callback    func(*config) error
}

// Declare the allCommands variable but assign the variables later to prevent an initialization cycle
var allCommands = map[string]cliCommands{}

func init() {
	allCommands = map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Lists all commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists a page of 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 locations",
			callback:    commandMapb,
		},
	}
}
