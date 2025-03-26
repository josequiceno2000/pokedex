package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

type cliCommand struct {
	name string
	description string
	callback func()
}

map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 
		
		words := cleanInput(scanner.Text())
		
		if len(words) == 0 {
			continue
		}

		
	}
}

func commandExit() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(trimmedText)
	words := strings.Fields(loweredText)
	return words
}