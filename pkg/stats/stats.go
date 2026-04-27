package stats

import "strings"

// CountVowels counts vowels in a string
func CountVowels(input string) int {
	vowels := "aeiouAEIOU"
	count := 0

	for _, ch := range input {
		if strings.ContainsRune(vowels, ch) {
			count++
		}
	}

	return count
}
