package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prompt = "Pokedex > "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		fmt.Println("Your command was:", cleanedInput[0])
	}
}
func cleanInput(text string) []string {
	split := strings.Fields(text)
	for i, word := range split {
		split[i] = strings.ToLower(word)
	}

	return split
}
