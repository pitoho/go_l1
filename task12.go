package main

import "fmt"

func tassk12() {
	// Исходная последовательность строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество для уникальных строк
	stringSet := make(map[string]bool)

	// Добавляем строки в множество
	for _, str := range sequence {
		stringSet[str] = true
	}

	// Выводим уникальные строки из множества
	fmt.Println("Уникальные строки в множестве:", getUniqueStrings(stringSet))
}

// Функция для получения уникальных строк из множества
func getUniqueStrings(set map[string]bool) []string {
	uniqueStrings := make([]string, 0, len(set))
	for str := range set {
		uniqueStrings = append(uniqueStrings, str)
	}
	return uniqueStrings
}
