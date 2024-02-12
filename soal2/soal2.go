package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortLetters(data []string) string {
	letterCounts := make(map[rune]int)

	for _, word := range data {
		for _, char := range word {
			letterCounts[char]++
		}
	}

	var pairs []Pair
	for char, count := range letterCounts {
		pairs = append(pairs, Pair{char, count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Count != pairs[j].Count {
			return pairs[i].Count > pairs[j].Count
		}
		return pairs[i].Char < pairs[j].Char
	})

	var result strings.Builder
	for _, pair := range pairs {
		for i := 0; i < pair.Count; i++ {
			result.WriteRune(pair.Char)
		}
	}

	return result.String()
}

type Pair struct {
	Char  rune
	Count int
}

func main() {
	data1 := []string{"Abc", "bCd"}
	data2 := []string{"Oke", "One"}
	data3 := []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}

	fmt.Println("Input:", data1)
	fmt.Println("Output:", sortLetters(data1))
	fmt.Println("Input:", data2)
	fmt.Println("Output:", sortLetters(data2))
	fmt.Println("Input:", data3)
	fmt.Println("Output:", sortLetters(data3))
}
