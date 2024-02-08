package main

import (
	"fmt"
	"sort"
)

func task16() {
	// Исходный массив
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}

	// Используем встроенную функцию сортировки
	quickSort(arr)

	// Выводим отсортированный массив
	fmt.Println("Отсортированный массив:", arr)
}

func quickSort(arr []int) {
	// Используем встроенную функцию сортировки
	sort.Ints(arr)
}
