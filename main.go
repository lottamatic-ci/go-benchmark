package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lottamatic-ci/go-benchmark/internal/textutils"
	"lottamatic-ci/go-benchmark/pkg/stats"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a sentence: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Use internal package
	wordCount := textutils.CountWords(input)

	// Use reusable package
	vowelCount := stats.CountVowels(input)

	fmt.Println("\n--- Results ---")
	fmt.Printf("Words: %d\n", wordCount)
	fmt.Printf("Vowels: %d\n", vowelCount)

	if wordCount > 5 {
		fmt.Println("That's a long sentence!")
	} else {
		fmt.Println("Nice and short.")
	}
}
