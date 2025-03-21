package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	cleanOutput := []string{}
	
	if len(text) > 0 {
		text = strings.Trim(text, " ")
		text = strings.ToLower(text)
		wordsToAdd := strings.Fields(text)
		cleanOutput = append(cleanOutput, wordsToAdd...)
	}
	return cleanOutput

}