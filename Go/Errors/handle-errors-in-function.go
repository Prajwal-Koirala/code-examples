package main

import (
	"log"
	"os"
)

func main() {
	// Handle all errors at the function level.
	err := os.MkdirAll("assets/remove/6VoNjRbnPXfmLpW1WyWt", 0755)
	if err != nil {
		log.Fatalln(err)
	}
}
