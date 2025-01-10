package main

import (
	"log"
	"os"
)

func createEmptyFile(filename string) {
	// Create the file with read-write permissions for the owner
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	defer file.Close() // Ensure the file is closed after creation

	// File is created empty, no need to write anything
	log.Println("File created successfully:", filename)
}

func main() {
	createEmptyFile("emptyfile.txt")
}
