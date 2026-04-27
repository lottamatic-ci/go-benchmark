package textutils

import "strings"

// CountWords returns the number of words in a string
func CountWords(input string) int {
	words := strings.Fields(input)
	return len(words)
}
