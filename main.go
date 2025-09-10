package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	text, _ = strings.CutPrefix(text, " ")
	text, _ = strings.CutSuffix(text, " ")
	text = strings.ToLower(text)

	words := strings.Split(text, " ")
	
	return words
}