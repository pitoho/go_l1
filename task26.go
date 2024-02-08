package main

import (
	"fmt"
	"strings"
)

func areAllCharactersUnique(s string) bool {
	// Приводим строку к нижнему регистру
	lowercaseString := strings.ToLower(s)

	// Создаем карту для отслеживания уникальных символов
	characterMap := make(map[rune]bool)

	// Перебираем каждый символ в строке
	for _, char := range lowercaseString {
		// Если символ уже есть в карте, значит, строка не уникальна
		if characterMap[char] {
			return false
		}
		// Добавляем символ в карту
		characterMap[char] = true
	}

	// Если дошли до этой точки, все символы уникальны
	return true
}

func task26() {
	fmt.Println(areAllCharactersUnique("abcd"))     
	fmt.Println(areAllCharactersUnique("abCdefAaf")) 
	fmt.Println(areAllCharactersUnique("aabcd"))    
}
