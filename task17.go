package main

import (
	"fmt"
	"sort"
)

func task17() {
	// Исходный отсортированный слайс
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	// Элемент, который мы ищем
	target := 12

	// Используем встроенный бинарный поиск
	index := binarySearch(arr, target)

	// Выводим результат
	if index != -1 {
		fmt.Printf("Элемент %d найден в позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в слайсе\n", target)
	}
}

func binarySearch(arr []int, target int) int {
	// Используем встроенный бинарный поиск
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	// Проверяем, был ли найден элемент
	if index < len(arr) && arr[index] == target {
		return index
	} else {
		return -1
	}
}
