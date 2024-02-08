package main

import "fmt"

func task11() {
	set1 := map[string]bool{
		"apple":  true,
		"banana": true,
		"orange": true,
		"grape":  true,
	}

	set2 := map[string]bool{
		"banana": true,
		"pear":   true,
		"kiwi":   true,
		"grape":  true,
	}

	// Получаем пересечение множеств
	intersection := intersectSets(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:", intersection)
}

// Функция для нахождения пересечения множеств
func intersectSets(set1, set2 map[string]bool) map[string]bool {
	intersection := make(map[string]bool)

	// Проходим по элементам первого множества
	for element := range set1 {
		// Если элемент есть во втором множестве, добавляем его в пересечение
		if set2[element] {
			intersection[element] = true
		}
	}

	return intersection
}
