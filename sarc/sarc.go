package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	text := strings.Trim(strings.Join(os.Args[1:], ""), " ")

	if len(text) == 0 {
		fmt.Println("USAGE: sarc The text you want to be sarcastic")

		os.Exit(1)
	}

	sarcasticText := ToSarcastic(text)
	fmt.Println(sarcasticText, time.Since(start))
}

func ToSarcastic(text string) string {
	output := []rune(text)
	shouldUpperCase := true

	for idx, character := range text {
		if unicode.IsLetter(character) {
			shouldUpperCase = !shouldUpperCase
		}

		output[idx] = unicode.ToLower(character)
		if shouldUpperCase {
			output[idx] = unicode.ToUpper(character)
		}
	}

	return string(output)
}
