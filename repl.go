package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

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

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(trimmedText)
	words := strings.Fields(loweredText)
	return words
}