package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prompt = "Pokedex > "

var configs = config{"", ""}

func cleanInput(text string) []string {
	split := strings.Fields(text)
	for i, word := range split {
		split[i] = strings.ToLower(word)
	}

	return split
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		firstWord := cleanedInput[0]
		command, exists := commands[firstWord]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&configs)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
