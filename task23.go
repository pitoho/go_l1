package main

import "fmt"

func removeElement(slice []int, index int) []int {
	// Проверяем, что индекс находится в пределах слайса
	if index < 0 || index >= len(slice) {
		return slice // Если индекс некорректен, возвращаем исходный слайс без изменений
	}

	// Используем срезы для удаления элемента по индексу
	return append(slice[:index], slice[index+1:]...)
}

func task23() {
	// Пример использования
	slice := []int{1, 2, 3, 4, 5}

	// Удаляем элемент с индексом 2
	indexToRemove := 2
	newSlice := removeElement(slice, indexToRemove)

	// Выводим результат
	fmt.Printf("Исходный слайс: %v\n", slice)
	fmt.Printf("Слайс после удаления элемента с индексом %d: %v\n", indexToRemove, newSlice)
}
