package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Конвертируем строку в слайс рун
	runes := []rune(s)

	// Получаем длину строки в байтах
	length := len(runes)

	// Индексы начала и конца строки
	start, end := 0, length-1

	// Переворачиваем строку в слайсе рун
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	// Конвертируем слайс рун обратно в строку
	return string(runes)
}

func task19() {
	// Пример использования
	inputString := "главрыба"
	reversedString := reverseString(inputString)

	// Выводим результат
	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Перевернутая строка: %s\n", reversedString)
}
