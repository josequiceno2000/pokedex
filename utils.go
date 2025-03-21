package main

import (
	"fmt"
	"strings"
	"os"
)

func cleanInput(text string) []string {
	text = strings.ToLower(strings.TrimSpace(text))
	return strings.Fields(text)
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}