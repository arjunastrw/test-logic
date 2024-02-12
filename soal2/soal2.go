package main

import (
	"fmt"
	"sort"
	"strings"
)

func printInput(strs []string) {
	fmt.Print("Input:[")
	for i, str := range strs {
		fmt.Printf(`"%s"`, str)
		if i < len(strs)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func charOrganizer(strs []string) {
	printInput(strs)

	str := strings.Join(strs, "")
	charMap := make(map[rune]int)
	for _, char := range str {
		charMap[char]++
	}

	var keys []rune
	values := make(map[int]int)
	for key, value := range charMap {
		keys = append(keys, key)
		values[value]++
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return charMap[keys[i]] > charMap[keys[j]]
	})

	var asciiSortStr []string
	fmt.Println("Output:")
	flag := 1
	for _, char := range keys {
		flag = values[charMap[char]]
		if flag == 1 {
			if len(asciiSortStr) != 0 {
				asciiSortStr = append(asciiSortStr, string(char))
				sort.SliceStable(asciiSortStr, func(i, j int) bool {
					return asciiSortStr[i] < asciiSortStr[j]
				})
				newStr := strings.Join(asciiSortStr, "")

				fmt.Print(newStr)
				asciiSortStr = nil
			} else {
				fmt.Print(string(char))
			}
		} else {
			asciiSortStr = append(asciiSortStr, string(char))
			values[charMap[char]]--
		}
	}
	fmt.Println()
}

func main() {
	var strs []string

	strs = []string{"Abc", "bCd"}
	charOrganizer(strs)

	strs = []string{"One", "Oke"}
	charOrganizer(strs)

	strs = []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}
	charOrganizer(strs)
}
