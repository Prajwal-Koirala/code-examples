package main

import (
	"fmt"
	"sort"
)

// sortMapByValues returns a sorted slice of key-value pairs from the input map based on values.
func sortMapByValues(inputMap map[string]string) [][]string {
	// Extract key-value pairs
	pairs := make([][2]string, 0, len(inputMap))
	for key, value := range inputMap {
		pairs = append(pairs, [2]string{key, value})
	}

	// Sort pairs by values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	// Convert to the expected return type
	result := make([][]string, len(pairs))
	for i, pair := range pairs {
		result[i] = []string{pair[0], pair[1]}
	}
	return result
}

func main() {
	keyValueMap := map[string]string{
		"b": "false",
		"a": "true",
		"d": "false",
		"c": "true",
	}

	// Print sorted key-value pairs by values
	for _, pair := range sortMapByValues(keyValueMap) {
		fmt.Println(pair[0] + ": " + pair[1])
	}
}
