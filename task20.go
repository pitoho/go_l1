package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func task20() {
	inputString := "snow dog sun"
	reversedString := reverseWords(inputString)

	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Строка с перевернутыми словами: %s\n", reversedString)
}
