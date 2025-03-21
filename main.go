package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	for _, command := range supportedCommands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

var supportedCommands map[string]cliCommand

func main()  {
	supportedCommands = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")

	for scanner.Scan() {
		userInput := scanner.Text()
		cleanedUserInput := cleanInput(userInput)
		userCommand := cleanedUserInput[0]


		command, ok := supportedCommands[userCommand]
		if ok {
			switch command.name {
			case "exit":
				err := command.callback()
				if err != nil {
					fmt.Errorf("Error: %w", err)
				}
				break;
			case "help":
				err := command.callback()
				if err != nil {
					fmt.Errorf("error: %w", err)
				}
				break;
			}
			
		} else {
			fmt.Errorf("Unknown command")
		}
		
		fmt.Print("Pokedex > ")
	}
}

