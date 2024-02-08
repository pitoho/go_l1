package main

import (
	"fmt"
	"math"
)

func task22() {
	// Задаем значения переменных a и b (больше чем 2^20)
	a := math.Pow(2, 21)
	b := math.Pow(2, 22)

	// Умножение
	resultMultiply := multiply(a, b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, resultMultiply)

	// Деление
	resultDivide, err := divide(a, b)
	if err != nil {
		fmt.Println("Ошибка при делении:", err)
	} else {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, resultDivide)
	}

	// Сложение
	resultAdd := add(a, b)
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, resultAdd)

	// Вычитание
	resultSubtract := subtract(a, b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, resultSubtract)
}

// Умножение
func multiply(a, b float64) float64 {
	return a * b
}

// Деление
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a / b, nil
}

// Сложение
func add(a, b float64) float64 {
	return a + b
}

// Вычитание
func subtract(a, b float64) float64 {
	return a - b
}
