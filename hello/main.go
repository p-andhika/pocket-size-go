package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur...")
	flag.Parse()

	greeting := greet(language(lang))
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
