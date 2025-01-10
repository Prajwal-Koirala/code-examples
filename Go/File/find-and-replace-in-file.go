package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Find and replace content in a given file.
	findAndReplaceInFile("example.md", "{REPLACE_CONTENT_WITH_GOLANG}", "Your new content goes here")
}

// Find a given content in a given file and replace it with given content.
func findAndReplaceInFile(filePath string, prefixContent string, givenContent string) {
	// Read the content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}
	// Convert content to string and replace the target string
	updatedContent := strings.ReplaceAll(string(content), prefixContent, givenContent)
	// Write the updated content back to the file
	err = ioutil.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		log.Println(err)
	}
}
