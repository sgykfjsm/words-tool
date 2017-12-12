package main

import (
	"fmt"
	"log"
	"os"

	wordsclient "github.com/sgykfjsm/words-tool/words"
)

func Usage() {
	fmt.Printf("Usage: %s word\n", os.Args[0])
}

func DebugPrint(result *wordsclient.WordsResponse) {
	for _, r := range result.Results {
		fmt.Println("-", r.Definition)
		for _, example := range r.Examples {
			fmt.Println("    -", example)
		}
	}
}

func main() {
	key := os.Getenv("MASHAPE_KEY")
	if key == "" {
		fmt.Println("Set your Mashape Key as environment varilas 'MASHAPE_KEY'")
		Usage()
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Please specify the word you want to search")
		Usage()
		os.Exit(1)
	}
	word := os.Args[1]

	words := wordsclient.New(key)
	result, err := words.Words(word)
	if err != nil {
		log.Fatal(err)
	}

	DebugPrint(result)
}
