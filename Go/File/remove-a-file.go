package main

import (
	"log"
	"os"
)

func main() {
	// Remove a file
	removeFile("bar.txt")
}

// Remove a file from the file system
func removeFile(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatalln(err)
	}
}