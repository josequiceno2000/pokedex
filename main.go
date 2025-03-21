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


func main()  {
	supportedCommands := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
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
			err := command.callback()
			if err != nil {
				fmt.Errorf("Error: %w", err)
			}
			break;
		} else {
			fmt.Errorf("Unknown command")
		}
		
		fmt.Print("Pokedex > ")
	}
}

