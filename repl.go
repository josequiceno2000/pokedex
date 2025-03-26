package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

func startRepl() {
	type cliCommand struct {
		name string
		description string
		callback func() error
	}
	
	possibleCommands := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 
		
		words := cleanInput(scanner.Text())
		
		if len(words) == 0 {
			continue
		} else {
			command := words[0]
			_, exists := possibleCommands[command]

			if exists {
				err := possibleCommands[command].callback
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}

		
	}
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(trimmedText)
	words := strings.Fields(loweredText)
	return words
}