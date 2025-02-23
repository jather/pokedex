package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	split := strings.Fields(text)
	for i, word := range split {
		split[i] = strings.ToLower(word)
	}

	return split
}
func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cleanedInput := cleanInput(input)
		firstWord := cleanedInput[0]
		argument := ""
		if len(cleanedInput) > 1 {
			argument = cleanedInput[1]
		}

		command, exists := commands[firstWord]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(config, argument)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
