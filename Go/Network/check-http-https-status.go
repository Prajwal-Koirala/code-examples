package main

import (
	"fmt"
	"log"
	"net/http"
)

// CheckWebsiteHTTPStatus checks if the given domain is reachable via both HTTP and HTTPS.
// It returns true if at least one of the protocols returns an HTTP 200 OK status.
func CheckWebsiteHTTPStatus(domain string) bool {
	// Create a slice containing both HTTP and HTTPS protocols for checking
	protocols := []string{"http://", "https://"}
	// Loop through each protocol to attempt a request
	for _, protocol := range protocols {
		// Make an HTTP GET request using the current protocol and the specified domain
		resp, err := http.Get(protocol + domain)
		// If there's no error, proceed to check the response
		if err == nil {
			// Ensure the response body is closed after we're done with it to prevent resource leaks
			defer resp.Body.Close()
			// Check if the response status code is 200 OK
			if resp.StatusCode == http.StatusOK {
				return true // Return true if the website is reachable
			}
		} else {
			// Log any errors encountered during the request, but continue checking the next protocol
			log.Println(err)
		}
	}
	// Return false if neither protocol succeeded
	return false
}

func main() {
	// Specify the domain to check; can be replaced with any valid domain
	domain := "example.com"
	// Call the function to check if the website is online
	isOnline := CheckWebsiteHTTPStatus(domain)
	// Output the result based on the response from the function
	if isOnline {
		// Inform the user that the website is online
		fmt.Printf("The website %s is online.\n", domain)
	} else {
		// Inform the user that the website is offline
		fmt.Printf("The website %s is offline.\n", domain)
	}
}
