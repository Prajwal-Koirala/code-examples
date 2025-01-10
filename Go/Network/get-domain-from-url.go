package main

import (
	"fmt"
	"log"
	"net/url"
)

// getDomainFromURL extracts the domain name from the given URL string.
func getDomainFromURL(givenURL string) string {
	// Parse the given URL string into a URL structure
	parsedUrl, err := url.Parse(givenURL)
	if err != nil {
		// Log an error message if parsing fails
		log.Println(err)
	}
	// Return the hostname (domain name) from the parsed URL
	return parsedUrl.Hostname()
}

func main() {
	// Define a URL string to extract the domain from
	urlContent := "https://fmovies.ps/"
	// Call getDomainFromURL and print the extracted domain name
	fmt.Println(getDomainFromURL(urlContent))
}
