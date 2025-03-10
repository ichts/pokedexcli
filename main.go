package main

import (
	"fmt"
	"strings"
)

func main() {
	cleanInput("Hello       from ichts")

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
