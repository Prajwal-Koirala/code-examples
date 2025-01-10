package main 

import (
	"fmt"
	"sort"
)

// sortMapByKeys returns a sorted slice of key-value pairs from the input map.
func sortMapByKeys(inputMap map[string]string) [][]string {
	// Extract keys and sort them
	keys := make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Use sorted keys to populate sorted pairs
	pairs := make([][]string, len(inputMap))
	for i, key := range keys {
		pairs[i] = []string{key, inputMap[key]}
	}
	return pairs
}

func main() {
	keyValueMap := map[string]string{ 
		"b": "false", 
		"a": "true",  
		"d": "false", 
		"c": "true",  
	}

	// Print sorted key-value pairs
	for _, pair := range sortMapByKeys(keyValueMap) {
		fmt.Println(pair[0] + ": " + pair[1])
	}
}
