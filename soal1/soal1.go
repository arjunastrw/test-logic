package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func countLetters(input string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range input {
		char = unicode.ToLower(char)
		if unicode.IsLetter(char) {
			counts[char]++
		}
	}
	return counts
}

func formatOutput(counts map[rune]int) string {
	var pairs []string
	for char, count := range counts {
		pairs = append(pairs, fmt.Sprintf("%c=%d", char, count))
	}
	sort.Strings(pairs)
	output := strings.Join(pairs, ", ")
	return output
}

func main() {
	input1 := "We Always Mekar"
	input2 := "coding is fun"

	counts1 := countLetters(input1)
	counts2 := countLetters(input2)

	fmt.Println("Input:", input1)
	fmt.Println("Output:", formatOutput(counts1))
	fmt.Println("Input:", input2)
	fmt.Println("Output:", formatOutput(counts2))
}
