package main

import "fmt"

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Println("Usage:")

	for _, command := range getCommands() {
		fmt.Printf("\n%s: %s\n", command.name, command.description)
	}
	return nil
}