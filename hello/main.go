package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// language represent the language's code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"en": "Hello world",
	"id": "Halo dunia",
	"fr": "Bonjour le monde",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
