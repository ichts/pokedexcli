package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *cfg) {

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		if text == "" {
			continue
		}

		cleanText := strings.Fields(strings.ToLower(text))
		command, exists := commands[cleanText[0]]
		if exists {
			if command.name == "explore" || command.name == "catch" {
				config.Flag = cleanText[1]
			}
			command.callback(config)
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

func cleanInput(text string) []string {
	cleanText := strings.Fields(strings.ToLower(text))

	for i, w := range cleanText {
		fmt.Println(cleanText[i])
		cleanText[i] = strings.TrimSpace(w)
		fmt.Println(cleanText[i])
	}

	return cleanText
}
