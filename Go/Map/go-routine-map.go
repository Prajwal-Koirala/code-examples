package main

import (
	"fmt"
	"sync"
)

// saveToMap saves a key-value pair into the provided sync.Map
func saveToMap(safeMap *sync.Map, key string, value interface{}) {
	// Store the key-value pair in the given map
	safeMap.Store(key, value)
}

// retrieveFromMap retrieves a value from the provided sync.Map for a given key.
// If the key does not exist, it returns nil.
func retrieveFromMap(safeMap *sync.Map, key string) interface{} {
	// Load fetches the value for the key; if the key doesn't exist, it returns nil
	value, _ := safeMap.Load(key)
	return value
}

func main() {
	// Create multiple sync.Map instances
	var map1 sync.Map
	var map2 sync.Map

	// Save key-value pairs into the first map (map1)
	saveToMap(&map1, "name", "John Doe")
	saveToMap(&map1, "age", 30)

	// Save key-value pairs into the second map (map2)
	saveToMap(&map2, "city", "New York")
	saveToMap(&map2, "country", "USA")

	// Retrieve values from map1
	value := retrieveFromMap(&map1, "name")
	if value == nil {
		fmt.Println("Map1 - Name not found")
	} else {
		fmt.Println("Map1 - Name:", value)
	}

	value = retrieveFromMap(&map1, "age")
	if value == nil {
		fmt.Println("Map1 - Age not found")
	} else {
		fmt.Println("Map1 - Age:", value)
	}

	// Retrieve values from map2
	value = retrieveFromMap(&map2, "city")
	if value == nil {
		fmt.Println("Map2 - City not found")
	} else {
		fmt.Println("Map2 - City:", value)
	}

	value = retrieveFromMap(&map2, "country")
	if value == nil {
		fmt.Println("Map2 - Country not found")
	} else {
		fmt.Println("Map2 - Country:", value)
	}

	// Attempt to retrieve a non-existent key
	value = retrieveFromMap(&map2, "continent")
	if value == nil {
		fmt.Println("Map2 - Continent not found")
	} else {
		fmt.Println("Map2 - Continent:", value)
	}
}
