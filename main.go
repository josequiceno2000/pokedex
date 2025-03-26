package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		
		if scanner.Scan() {
			userInput := scanner.Text()
			cleanedInput := cleanInput(userInput)
			fmt.Printf("Your command was: %v\n", cleanedInput[0])
		}
	}
	
}