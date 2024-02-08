package main

import (
	"fmt"
	"math"
	"sort"
)

func task10() {
	// Исходная последовательность температур
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Группируем температуры по шагу 10 градусов
	groupedTemperatures := groupTemperatures(temperatures, 10)

	// Выводим результат
	for key, values := range groupedTemperatures {
		fmt.Printf("%2.1f to %2.1f: %v\n", key, key+10, values)
	}
}

// Функция для группировки температур по шагу
func groupTemperatures(temperatures []float64, step float64) map[float64][]float64 {
	groupedTemperatures := make(map[float64][]float64)

	// Сортируем температуры для упрощения группировки
	sort.Float64s(temperatures)

	// Группируем температуры по шагу
	for _, temp := range temperatures {
		key := roundToStep(temp, step)
		groupedTemperatures[key] = append(groupedTemperatures[key], temp)
	}

	return groupedTemperatures
}

// Функция для округления числа до ближайшего шага
func roundToStep(value, step float64) float64 {
	return math.Round(value/step) * step
}
